import { chatTemplate } from "../templates/chatTemplate.js";
import { getUsers } from "../services/chatService.js";
import { render } from "../utils/render.js";
import { connectWebSocket, on } from "../websocket/socket.js";

import { renderChatList, initChatList } from "../components/chatList.js";

import { initChatWindow, openConversation, addMessage, showTyping, hideTyping } from "../components/chatWindow.js";



let users = [];

let authUser = null;

let selectedUser = null;



export async function chatPage(user) {

    authUser = user;


    render(
        chatTemplate()
    );


    connectWebSocket();
   

    initChatWindow(
        user.id
    );


    await loadUsers();


    setupSocketEvents();

}



async function loadUsers() {

    try {

        users = await getUsers();


        const container =
            document.querySelector(
                "#chat-users-list"
            );


        renderChatList(
            container,
            users
        );


        initChatList(
            selectUser
        );


    } catch(err) {

        console.error(
            "Load users error:",
            err
        );

    }

}



async function selectUser(userID) {


    const user =
        users.find(
            u => u.id === userID
        );


    if (!user) return;


    selectedUser = user;


    await openConversation(
        user
    );

}



function setupSocketEvents() {


    on(
        "message",
        (message)=>{


            if (
                (message.sender_id === selectedUser.id && message.receiver_id === authUser.id) ||
                (message.sender_id === authUser.id && message.receiver_id === selectedUser.id)
            ){
                addMessage(message);
            }

        }
    );



    on(
        "typing",
        (data)=>{


            if (selectedUser && data.sender_id === selectedUser.id){

                if(data.is_typing){

                    showTyping(selectedUser.nickname);

                }else{

                    hideTyping();

                }

            }

        }
    );



    on(
        "online",
        (data)=>{


            const user =
                users.find(
                    u => u.id === data.user_id
                );


            if(user){

                user.is_online =
                    data.online;


                const container =
                    document.querySelector(
                        "#chat-users-list"
                    );


                renderChatList(
                    container,
                    users
                );


                initChatList(
                    selectUser
                );

            }

        }
    );


}
