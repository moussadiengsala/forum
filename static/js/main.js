const menuButton = document.querySelector("#menu-button");
const menuModal = document.querySelector("#menu-modal");
const menuClose = document.querySelector("#menu-close");

let toggleMenu = () => {
    menuModal.classList.toggle("-left-full");
    menuModal.classList.toggle("left-0");
};

menuButton.addEventListener("click", toggleMenu);
menuClose.addEventListener("click", toggleMenu);
