import { homePage } from "./pages/home.js";
import { loginPage } from "./pages/login.js";
import { registerPage } from "./pages/register.js";
import { postDetailsPage } from "./pages/postDetailsPage.js";
import { checkAuth, getMe } from "./services/authService.js";
import { errorPage } from "./pages/error.js";
import { chatPage } from "./pages/chat.js";
import ws from "./websocket/ws.js";
import { showRichNotification } from "./components/notification.js";
import { getCurrentReceiverID } from "./components/chatWindow.js";
import { escapeHTML } from "./utils/escapeHTML.js";


const routes = {
    "/": homePage,
    "/login": loginPage,
    "/register": registerPage,
    "/chat": chatPage,
};

export async function router() {

    const path = window.location.pathname;

    const isAuthenticated = await checkAuth();

    // User not logged in
    if (!isAuthenticated && path !== "/login" && path !== "/register") {
        navigate("/login");
        return;
    }

    if(isAuthenticated){

        ws.connect();

        // Cache current user ID for notification logic.
        // Must resolve BEFORE registering the listener to avoid
        // dropping the first incoming message.
        try {
            const user = await getMe();
            window.__currentUserId = user.id;
        } catch (e) {
            window.__currentUserId = null;
        }

        // Global listener for incoming chat messages
        setupChatNotifications();

    }
    // User already logged in
    if (isAuthenticated && (path === "/login" || path === "/register")) {

        navigate("/");
        return;
    }

    // Dynamic Route: /posts/:id
    if (path.startsWith("/posts/")) {

        const id = path.split("/")[2];

        if (!id || isNaN(id)) {
            errorPage();
            return;
        }

        postDetailsPage(Number(id));
        return;
    }

    const page = routes[path];

    if (!page) {
        errorPage();
        return;
    }

    page();

}

/**
 * Requests browser notification permission (called once on user interaction).
 */
function requestNotificationPermission() {
    if (!("Notification" in window)) return false;
    if (Notification.permission === "granted") return true;
    if (Notification.permission === "denied") return false;
    Notification.requestPermission();
    return false;
}

/**
 * Shows a browser notification (HTML5 Notification API).
 */
function sendBrowserNotification(title, body, senderId) {
    if (!("Notification" in window)) return false;
    if (Notification.permission !== "granted") return false;

    try {
        const notif = new Notification(title, {
            body,
            icon: "/favicon.ico",
            tag: "chat-message",
        });

        notif.onclick = () => {
            window.focus();
            window.__pendingChatUserId = senderId;
            navigate("/chat");
            notif.close();
        };

        return true;
    } catch (e) {
        return false;
    }
}

let chatNotificationsSetup = false;

/**
 * Sets up the global WebSocket listener for incoming chat messages.
 *
 * - If the user is already viewing that conversation → no notification.
 * - If the browser tab is hidden → tries the Browser Notification API first,
 *   then falls back to an in-app notification.
 * - Otherwise → shows a clickable in-app notification.
 *
 * Uses a guard flag so the listener is registered only once across navigations.
 */
function setupChatNotifications() {

    if (chatNotificationsSetup) return;
    chatNotificationsSetup = true;

    ws.on("message", (message) => {

        const myId = window.__currentUserId;

        // Ignore messages we sent ourselves
        if (!myId || message.sender_id === myId) return;

        // Check if currently viewing this conversation
        const currentPartnerID = getCurrentReceiverID();
        const isOnChatPage = window.location.pathname === "/chat";

        if (isOnChatPage && currentPartnerID === message.sender_id) {
            return; // Chat page already handles displaying the message
        }

        const senderName = escapeHTML(message.sender_nickname || "Someone");
        const content = escapeHTML(message.content || "");
        const senderId = message.sender_id;

        // Attempt browser notification if tab is hidden
        if (document.hidden) {
            requestNotificationPermission();
            const browserSent = sendBrowserNotification(
                `New message from ${senderName}`,
                content.substring(0, 120),
                senderId
            );
            if (browserSent) return;
        }

        // Fallback: in-app notification
        showRichNotification({
            html: `
                <div class="chat-notification-content">
                    <strong>${senderName}</strong>
                    <p>${content.substring(0, 100)}</p>
                </div>
            `,
            type: "chat",
            duration: 4000,
            onClick: () => {
                window.__pendingChatUserId = senderId;
                navigate("/chat");
            },
        });

    });

}


export function navigate(path) {
    window.history.pushState({}, "", path);
    router();
}

window.addEventListener("popstate", router);