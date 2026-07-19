import { getConversation, sendMessage, markConversationAsRead, sendTyping } from "../services/chatService.js";
import { messageBubble } from "./messageBubble.js";
import {ws} from "../websocket/ws.js";


let currentReceiverID = null;
let currentUserID = null;


export function initChatWindow() {

    const backBtn = document.querySelector("#chat-back-btn");


    if (backBtn) {

        backBtn.addEventListener("click", () => {

            document
                .querySelector(".chat-page")
                ?.classList.remove("show-chat");

        });

    }

    const form = document.querySelector("#chat-form");
    const input = document.querySelector("#chat-input");


    form.addEventListener("submit", (e) => {

        e.preventDefault();

        if (!currentReceiverID) return;

        const content = input.value.trim();

        if (!content) return;


        sendMessage(
            currentReceiverID,
            content
        );


        input.value = "";

    });


    input.addEventListener("input", () => {

        if (!currentReceiverID) return;


        sendTyping(
            currentReceiverID,
            true
        );


        clearTimeout(input.typingTimer);


        input.typingTimer = setTimeout(() => {

            sendTyping(
                currentReceiverID,
                false
            );

        }, 800);

    });

}



export async function openConversation(user) {

    currentReceiverID = user.id;


    document.querySelector("#chat-user-name")
        .textContent = user.nickname;


    updateUserStatus(user.is_online);



    const messages = await getConversation(user.id);

    ws.sendRead(user.id);

    renderMessages(messages);



    await markConversationAsRead(
        user.id
    );


    scrollToBottom();

}



export function renderMessages(messages = []) {

    const container = document.querySelector("#chat-messages");


    container.innerHTML = "";


    messages.forEach(message => {

        container.innerHTML += messageBubble(
            message,
            currentUserID
        );

    });


}



export function addMessage(message) {

    const container = document.querySelector("#chat-messages");


    container.innerHTML += messageBubble(
        message,
        currentUserID
    );


    scrollToBottom();

}



export function showTyping(username) {

    const indicator = document.querySelector(
        "#typing-indicator"
    );


    indicator.textContent =
        `${username} is typing...`;

}



export function hideTyping() {

    const indicator = document.querySelector(
        "#typing-indicator"
    );


    indicator.textContent = "";

}



function updateUserStatus(isOnline) {

    const status =
        document.querySelector("#chat-user-status");


    if (!status) return;


    if (isOnline) {

        status.classList.remove("offline");
        status.classList.add("online");

    } else {

        status.classList.remove("online");
        status.classList.add("offline");

    }

}



function scrollToBottom() {

    const container =
        document.querySelector("#chat-messages");


    if (!container) return;


    container.scrollTop =
        container.scrollHeight;

}

export function setCurrentUser(userID) {
    currentUserID = userID;
}


export function updateReadReceipts() {

    const statuses = document.querySelectorAll(".message-status");

    statuses.forEach(status => {
        status.textContent = "✓✓";
    });

}