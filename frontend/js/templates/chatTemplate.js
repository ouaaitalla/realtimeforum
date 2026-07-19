



export function chatTemplate() {
    return `
        <div class="chat-page">

            <aside class="chat-sidebar">

                <div class="chat-sidebar-header">
                    <h2>Messages</h2>
                </div>

                <div
                    id="chat-users-list"
                    class="chat-users-list"
                >
                </div>

            </aside>

            <section class="chat-content">

                <div
                    id="chat-header"
                    class="chat-header"
                >

                <button id="chat-back-btn" class="chat-back-btn"> ← </button>

                    <div class="chat-user-info">

                        <div
                            id="chat-user-status"
                            class="chat-user-status offline"
                        ></div>

                        <span id="chat-user-name">
                            Select a conversation
                        </span>

                    </div>

                </div>

                <div
                    id="chat-messages"
                    class="chat-messages"
                >

                    <div class="chat-empty-state">
                        Select a conversation to start chatting.
                    </div>

                </div>

                <div class="chat-typing-container">

                    <span id="typing-indicator"></span>

                </div>

                <form
                    id="chat-form"
                    class="chat-form"
                >

                    <textarea
                        id="chat-input"
                        rows="2"
                        maxlength="1000"
                        placeholder="Type a message..."
                    ></textarea>

                    <button
                        id="send-message-btn"
                        type="submit"
                    >
                        Send
                    </button>

                </form>

            </section>

        </div>
    `;
}