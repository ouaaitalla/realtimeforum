import { chatTemplate } from "../templates/chatTemplate.js";
import { getUsers } from "../services/chatService.js";
import { render } from "../utils/render.js";
import ws from "../websocket/ws.js";
import { renderChatList, initChatList } from "../components/chatList.js";
import { initChatWindow, openConversation, addMessage, showTyping, hideTyping, setCurrentUser} from "../components/chatWindow.js";
import { getMe } from "../services/authService.js";
import { appLayout } from "../layouts/appLayout.js";
import { initNavbar } from "../components/navbar.js";


let users = [];

let authUser = null;

let selectedUser = null;



export async function chatPage() {

    authUser = await getMe();

    setCurrentUser(authUser.id);

    render(appLayout(chatTemplate()));

    initNavbar();

    ws.connect();

    initChatWindow();

    await loadUsers();

    setupSocketEvents();

    // If navigated from a notification click, auto-open that conversation
    const pendingUserId = window.__pendingChatUserId;
    if (pendingUserId) {
        window.__pendingChatUserId = null;
        const user = users.find(u => u.id === pendingUserId);
        if (user) {
            // Small delay to let DOM settle, then open and show chat on mobile
            setTimeout(() => {
                selectedUser = user;
                openConversation(user);
                document.querySelector(".chat-page")?.classList.add("show-chat");
            }, 100);
        }
    }

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


    ws.on(
        "message",
        (message)=>{
            const otherUserId = message.sender_id === authUser.id? message.receiver_id : message.sender_id;

            const user = users.find(u => u.id === otherUserId);

            if (user) {
                user.last_message = message.content;
                user.last_message_time = message.created_at;

                const container = document.querySelector("#chat-users-list");

                renderChatList(
                    container,
                    users,
                    selectedUser?.id
                );

                initChatList(selectUser);
            }

            if (
                selectedUser &&
                (message.sender_id === selectedUser.id && message.receiver_id === authUser.id) ||
                (message.sender_id === authUser.id && message.receiver_id === selectedUser.id)
            ){
                addMessage(message);
            }

        }
    );



    ws.on(
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



    ws.on("online", (onlineUsers) => {

        users.forEach(user => {
            user.is_online = onlineUsers.includes(user.id);
        });

        const container = document.querySelector("#chat-users-list");

        renderChatList(container, users);

        initChatList(selectUser);

    });




}
