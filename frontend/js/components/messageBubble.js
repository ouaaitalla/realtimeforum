export function messageBubble(message, currentUserID) {

    const isMine = message.sender_id === currentUserID;

    return `
        <div class="message-wrapper ${isMine ? "mine" : "other"}">

            <div class="message-bubble">

                <p class="message-content">
                    ${escapeHTML(message.content)}
                </p>

                <div class="message-info">

                    <span class="message-time">
                        ${formatMessageTime(message.created_at)}
                    </span>

                    ${
                        isMine
                        ?
                        `
                        <span class="message-status">
                            ${message.is_read ? "✓✓" : "✓"}
                        </span>
                        `
                        :
                        ""
                    }

                </div>

            </div>

        </div>
    `;
}


function formatMessageTime(date) {

    if (!date) return "";

    const time = new Date(date);

    return time.toLocaleTimeString([], {
        hour: "2-digit",
        minute: "2-digit"
    });
}


function escapeHTML(text) {

    const div = document.createElement("div");

    div.textContent = text;

    return div.innerHTML;
}