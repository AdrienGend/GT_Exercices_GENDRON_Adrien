var cards = document.querySelectorAll(".card");

cards.forEach(function(card) {
  card.addEventListener("click", function() {
    cards.forEach(function(card) {
      card.classList.remove("active");
    });
    this.classList.add("active");
  });
});
