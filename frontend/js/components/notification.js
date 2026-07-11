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
