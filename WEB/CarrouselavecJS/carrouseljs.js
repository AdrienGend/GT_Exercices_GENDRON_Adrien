var slideIndex = 1;
showSlide(slideIndex);

function changeSlide(n) {
	showSlide(slideIndex += n);
}

function goToSlide(n) {
	showSlide(slideIndex = n);
}

function showSlide(n) {
	var i;
	var slides = document.getElementsByClassName("slide");
	var navItems = document.getElementsByClassName("nav-item");
	
	if (n > slides.length) {
		slideIndex = 1;
	}
	
	if (n < 1) {
		slideIndex = slides.length;
	}

	
	for (i = 0; i < slides.length; i++) {
		slides[i].style.display = "none";
	}

	for (i = 0; i < navItems.length; i++) {
		navItems[i].classList.remove("active");
	}

	
	slides[slideIndex-1].style.display = "block";
	navItems[slideIndex-1].classList.add("active");
}
