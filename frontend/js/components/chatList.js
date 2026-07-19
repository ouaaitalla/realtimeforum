export function chatList(users = [], selectedUserId = null) {
    if (!users.length) {
        return `
            <div class="chat-users-empty">
                No users found.
            </div>
        `;
    }

    return users.map(user => `
        <div
            class="chat-user-card ${selectedUserId === user.id ? "active" : ""}"
            data-user-id="${user.id}"
        >

            <div class="chat-user-avatar">

                ${
                    user.avatar
                        ? `<img src="${user.avatar}" alt="${user.nickname}">`
                        : `<span>${user.nickname.charAt(0).toUpperCase()}</span>`
                }

                <span class="online-indicator ${user.is_online ? "online" : "offline"}"></span>

            </div>

            <div class="chat-user-content">

                <div class="chat-user-top">

                    <span class="chat-user-name">
                        ${user.nickname}
                    </span>

                </div>

                <div class="chat-user-bottom">

                    <span class="chat-last-message">
                        ${user.last_message || "No messages yet"}
                    </span>

                </div>

            </div>

        </div>
    `).join("");
}

export function renderChatList(container, users, selectedUserId = null) {
    container.innerHTML = chatList(users, selectedUserId);
}

export function initChatList(onSelectUser) {

    document
        .querySelectorAll(".chat-user-card")
        .forEach(card => {

            card.addEventListener("click", () => {

                const userId = Number(card.dataset.userId);


                // Mobile: hide users list and show chat
                document
                    .querySelector(".chat-page")
                    ?.classList.add("show-chat");


                onSelectUser(userId);

            });

        });

}

