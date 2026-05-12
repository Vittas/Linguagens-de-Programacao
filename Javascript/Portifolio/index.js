
function changeColor(cor){
    elements = [...document.getElementsByClassName("main")]

    elements.map((e) => {
        e.style.backgroundColor = cor
    })
    // let Element = document.getElementsByClassName("main");
    // let color = document.getElementsByClassName("blue").style.backgroundColor;
    // Element.style.backgroundColor = color
    
}