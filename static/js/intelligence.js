document.addEventListener("DOMContentLoaded", () => {

    const tabs =
        document.querySelectorAll(
            ".intelligence-tab"
        );

    const experiences =
        document.querySelectorAll(
            ".intelligence-experience"
        );


    if (!tabs.length || !experiences.length) {
        return;
    }


    tabs.forEach((tab) => {

        tab.addEventListener("click", () => {

            const selectedRole =
                tab.dataset.role;


            /*
             * Update tabs
             */

            tabs.forEach((item) => {

                const isActive =
                    item === tab;

                item.classList.toggle(
                    "active",
                    isActive
                );

                item.setAttribute(
                    "aria-selected",
                    String(isActive)
                );

            });


            /*
             * Update experience
             */

            experiences.forEach((experience) => {

                const shouldShow =
                    experience.dataset.experience ===
                    selectedRole;

                experience.classList.toggle(
                    "active",
                    shouldShow
                );

            });

        });

    });

});