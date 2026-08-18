document.addEventListener("DOMContentLoaded", () => {

    const generateButton =
        document.querySelector(".quiz-generate");

    const status =
        document.querySelector(".ai-status");


    if (!generateButton || !status) {
        return;
    }


    generateButton.addEventListener("click", () => {

        const originalText =
            generateButton.innerHTML;


        generateButton.innerHTML =
            "<span>✦</span> Generating...";

        generateButton.disabled = true;

        status.textContent =
            "Generating";


        setTimeout(() => {

            generateButton.innerHTML =
                originalText;

            generateButton.disabled =
                false;

            status.textContent =
                "Ready";

        }, 1400);

    });

});