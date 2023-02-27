// Sticky navbar header
window.onscroll = function() {stickyHeader()};

var header = document.querySelector("header");

var sticky = header.offsetTop;

function stickyHeader() {
    if (window.pageYOffset > sticky) {
        header.classList.add("sticky");
    } else {
        header.classList.remove("sticky");
    }
}

// Sticky vertical navbar
window.onscroll = function() {stickyNav()};

var nav = document.querySelector("nav");

var stickyNav = nav.offsetTop + 50;

function stickyNav() {
    if (window.pageYOffset > stickyNav) {
        nav.classList.add("sticky-vertical");
    } else {
        nav.classList.remove("sticky-vertical");
    }
}

// Carousel
var carouselIndex = 1;

showCarousel(carouselIndex);

function changeCarousel(n) {
    showCarousel(carouselIndex += n);
}

function showCarousel(n) {
    var i;
    var carouselItems = document.querySelectorAll(".carousel li");
    if (n > carouselItems.length) {carouselIndex = 1}
    if (n < 1) {carouselIndex = carouselItems.length}
    for (i = 0; i < carouselItems.length; i++) {
        carouselItems[i].style.display = "none";
    }
    carouselItems[carouselIndex-1].style.display = "block";
}
