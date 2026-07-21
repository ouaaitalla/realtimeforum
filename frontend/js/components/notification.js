let notificationContainer = null;

function createNotificationContainer() {

    if (notificationContainer) return;

    notificationContainer = document.createElement("div");

    notificationContainer.id = "notification-container";

    document.body.appendChild(notificationContainer);
}


export function showNotification(message, type = "success") {

    createNotificationContainer();


    const notification = document.createElement("div");

    notification.className = `notification ${type}`;

    notification.textContent = message;


    notificationContainer.appendChild(notification);


    setTimeout(() => {
        notification.classList.add("show");
    }, 10);


    setTimeout(() => {

        notification.classList.remove("show");

        setTimeout(() => {
            notification.remove();
        }, 300);

    }, 1000);
}


/**
 * Shows a rich, clickable notification with HTML content.
 *
 * @param {Object} options
 * @param {string} options.html        - Inner HTML of the notification
 * @param {string} [options.type]      - CSS class (e.g. "success", "error", "chat")
 * @param {number} [options.duration]  - Auto-dismiss time in ms (default 4000)
 * @param {Function} [options.onClick] - Called when the notification is clicked
 */
export function showRichNotification({ html, type = "chat", duration = 4000, onClick }) {

    createNotificationContainer();

    const notification = document.createElement("div");

    notification.className = `notification ${type}`;

    notification.innerHTML = html;

    if (onClick) {
        notification.style.cursor = "pointer";
        notification.addEventListener("click", (e) => {
            onClick(e);
            notification.classList.remove("show");
            setTimeout(() => notification.remove(), 300);
        });
    }

    notificationContainer.appendChild(notification);

    setTimeout(() => {
        notification.classList.add("show");
    }, 10);

    setTimeout(() => {
        notification.classList.remove("show");
        setTimeout(() => notification.remove(), 300);
    }, duration);

    return notification;
}
