

export function openModal(content, onOpen = null) {

    let modal = document.getElementById("modal");

    if (!modal) {
        modal = document.createElement("div");
        modal.id = "modal";
        modal.className = "modal-overlay";

        document.body.appendChild(modal);
    }

    modal.innerHTML = `
        <div class="modal-content">

            <button class="modal-close" id="close-modal">
                &times;
            </button>

            ${content}

        </div>
    `;

    modal.classList.add("show");

    document
        .getElementById("close-modal")
        .addEventListener("click", closeModal);

    modal.addEventListener("click", (event) => {
        if (event.target === modal) {
            closeModal();
        }
    });

    if (typeof onOpen === "function") {
        onOpen();
    }
}

export function closeModal() {

    const modal = document.getElementById("modal");

    if (!modal) return;

    modal.remove();
}
