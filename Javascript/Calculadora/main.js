let Equacao =  document.getElementById("equacao");
let Calcular = document.getElementById("calcular");

Calcular.addEventListener('click', FazerCalculo);

function FazerCalculo(){
    let valores =  Equacao.value;
    console.log(valores)

    let num = 0;

    for (i of valores){


        let num2 = parseInt(i);
        num = num + num2
        console.log(num)

        
        /// Sinais de mais e menos //////

        if(i === "+" || i === "-" || i === "*" || i === "/"){
            let sinais_basicos = i;
            console.log("este são os sinais basicos: " + sinais_basicos);

        };


    };

    // Equacao.value = eval(Equacao.value)
    
};