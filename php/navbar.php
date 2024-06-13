<?php include "conexaobanco.php"; ?>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="CSS/navbar.css">
    <title>Document</title>
    <nav>
        <?php
        echo $_SESSION['nome'];
        ?>
        <button id="sair">Sair</button> 
        <div id="teste">

        </div>

        <script>            
            let sairconta = document.getElementById("sair");
            sairconta.addEventListener('click', sair)
            function sair(){
                document.location.href = 'http://localhost/site_1/login.php'
            }
        </script>


    </nav>