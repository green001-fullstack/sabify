document.addEventListener("DOMContentLoaded", () => {

    const animatedElements =
        document.querySelectorAll(
            ".fade-up, .learning-step"
        );


    if (!animatedElements.length) {
        return;
    }


    const observer =
        new IntersectionObserver(
            (entries, observer) => {

                entries.forEach((entry) => {

                    if (!entry.isIntersecting) {
                        return;
                    }


                    entry.target.classList.add(
                        "is-visible"
                    );


                    observer.unobserve(
                        entry.target
                    );

                });

            },
            {
                threshold: 0.12
            }
        );


    animatedElements.forEach((element) => {

        observer.observe(element);

    });

});