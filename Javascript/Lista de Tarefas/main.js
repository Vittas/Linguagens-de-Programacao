let CriarTarefas  = document.getElementById("criar_tarefas");
let FormTarefa = document.getElementById("form_tarefa");
let ValueTarefa = document.getElementById("value_tarefa")
let AdicionarTarefa = document.getElementById("adicionar_tarefa");
let ListaTarefas = document.getElementById("lista_tarefas");

CriarTarefas.addEventListener('click',Abrirtarefa);
AdicionarTarefa.addEventListener('click', AdicionarTarefas)
ValueTarefa.addEventListener('key', (e) => {
    if (e.key === "Enter") {
      AdicionarTarefas();
    }
  });

function Abrirtarefa(){
    if (FormTarefa.style.visibility==='hidden'){
        FormTarefa.style.visibility='visible';
    }
    else{
        FormTarefa.style.visibility='hidden';

    }
};


function AdicionarTarefas(){
    let TarefaTexto = ValueTarefa.value;
    if (TarefaTexto.trim !== ""){
        const NovaTarefa = document.createElement("li");
        NovaTarefa.innerHTML = `${TarefaTexto} <button class="delete_button">Excluir</button>`;
        ListaTarefas.appendChild(NovaTarefa);

        

    }
}

ListaTarefas.addEventListener('click', function(e){
    if (e.target.classList.contains("delete_button")){
        e.target.parentElement.remove();
    }
})



// function CriarTarefa(){]
//     const novaTarefa =  document.createElement("div")
//     novaTarefa.innerHTML = '<div class="model"><input id="tarefa" type="text" placeholder="Adicione uma tarefa"> <button id="adicionar">Adicionar</button></div>'
//     mostrar_tarefas.appendChild(novaTarefa)
    
//   }


///////////////////// funcionando //////////////////////
// function Abrirtarefa(){
//     const FormTarefa = document.getElementById("form_tarefa").style.visibility='visible';
// }







//////////////// Função de gerar input de tarefas //////////////







// const input_tarefas = document.getElementById("tarefa")
// const btn_add_task = document.getElementById("adicionar")
// const mostrar_tarefas = document.getElementById("tarefas")

// btn_add_task.addEventListener("click", adicionar_tarefa);
// input_tarefas.addEventListener("keypress", (e) => {
//   if (e.key === "Enter") {
//     adicionar_tarefa();
//   }
// });

// const adicionar_tarefa = () => {
//   const tarefaTexto = input_tarefas.value;
//   if (tarefaTexto.trim() !== "") {
//     const novaTarefa = document.createElement("li");
//     novaTarefa.innerHTML = `${tarefaTexto} <button class="excluir"> Excluir </button>`;
//     mostrar_tarefas.appendChild(novaTarefa)

//   }
// }

// mostrar_tarefas.addEventListener("click", function (e) {
//   if (e.target.classList.contains("Excluir")) {
//     e.target.parentElement.remove();
//   }
// });



