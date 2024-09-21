document.addEventListener("DOMContentLoaded", function() {
    const elements = document.querySelectorAll('.hidden-element');
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('visible');
            }
        });
    });

    elements.forEach(element => {
        observer.observe(element);
    });
});


let correo = document.querySelector(".correo")

let botonContacto = document.getElementById("botonContacto").addEventListener("click", function(){
    Swal.fire({
        title: "Contacto",
        html: correo,
        icon: "info"
    });
})


document.getElementById('toggleButton').addEventListener('click', function() {
    var menu = document.getElementById('dropdownMenu');
    menu.classList.toggle('hidden');
});


let botonSonido = document.querySelector(".botonSonido")

let botones = document.querySelectorAll("button")

for (i=0;i<botones.length;i++){
    botones[i].addEventListener("click", function(){
    botonSonido.play()
})
}
