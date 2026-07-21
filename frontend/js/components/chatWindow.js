import { getConversation, sendMessage, sendTyping } from "../services/chatService.js";
import { messageBubble } from "./messageBubble.js";
import { debounce } from "../utils/debounce.js";

const PAGE_SIZE = 10;

let currentReceiverID = null;
let currentUserID = null;
let currentOffset = 0;
let hasMoreMessages = true;
let isLoadingOlder = false;

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

    // Setup scroll-to-top pagination with debouncing
    const messagesContainer = document.querySelector("#chat-messages");

    if (messagesContainer) {

        const handleScroll = debounce(async () => {

            if (isLoadingOlder || !hasMoreMessages) return;

            // Detect when user scrolls near the top (within 50px)
            if (messagesContainer.scrollTop <= 50) {
                await loadOlderMessages();
            }

        }, 200);

        messagesContainer.addEventListener("scroll", handleScroll);

    }

}



export async function openConversation(user) {

    currentReceiverID = user.id;

    // Reset pagination state
    currentOffset = 0;
    hasMoreMessages = true;
    isLoadingOlder = false;


    document.querySelector("#chat-user-name")
        .textContent = user.nickname;


    updateUserStatus(user.is_online);


    // Load initial page (last 10 messages in DESC order from backend)
    const messages = await getConversation(user.id, PAGE_SIZE, 0);

    currentOffset = messages.length;

    if (messages.length < PAGE_SIZE) {
        hasMoreMessages = false;
    }

    // Backend returns DESC (newest first). Reverse for display (oldest first, newest at bottom)
    renderMessages(messages.reverse());


    scrollToBottom();

}


async function loadOlderMessages() {

    if (isLoadingOlder || !hasMoreMessages || !currentReceiverID) return;

    isLoadingOlder = true;

    const container = document.querySelector("#chat-messages");

    if (!container) {
        isLoadingOlder = false;
        return;
    }

    // Create and insert loading spinner at the top of the messages container
    const loadingIndicator = document.createElement("div");
    loadingIndicator.className = "chat-loading-older";
    loadingIndicator.id = "chat-loading-older";
    loadingIndicator.innerHTML = `
        <div class="chat-spinner"></div>
        <span>Loading older messages...</span>
    `;
    container.insertBefore(loadingIndicator, container.firstChild);

    // Save scroll height before adding messages
    const previousScrollHeight = container.scrollHeight;

    try {

        const messages = await getConversation(currentReceiverID, PAGE_SIZE, currentOffset);

        if (messages.length < PAGE_SIZE) {
            hasMoreMessages = false;
        }

        currentOffset += messages.length;

        if (messages.length === 0) {
            isLoadingOlder = false;
            return;
        }

        // Backend returns DESC (newest first). Reverse so oldest-first for prepending
        const reversed = messages.reverse();

        // Prepend older messages (insert before the first child)
        reversed.forEach(message => {
            container.insertAdjacentHTML(
                "afterbegin",
                messageBubble(message, currentUserID)
            );
        });

        // Maintain scroll position after prepending
        container.scrollTop = container.scrollHeight - previousScrollHeight;

    } catch (err) {
        console.error("Failed to load older messages:", err);
    } finally {
        isLoadingOlder = false;

        // Remove the loading spinner
        const spinner = document.querySelector("#chat-loading-older");
        if (spinner) {
            spinner.remove();
        }
    }

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


