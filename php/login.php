<?php include "conexaobanco.php"; ?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <nav>Login</nav>
    <form method="post">
        <h1>Logar</h1>
        <label for="nome">nome: </label>
        <input type="text" name="nome">
        <label for="senha">senha: </label>
        <input type="text" name="senha">
        <input type="submit" name="submit" value="submit">
    </form>

    <?php
        if(isset($_POST['submit'])){
            $nome = $_POST['nome'];
            $senha = $_POST['senha'];

            $SEARCH_SQL = "SELECT nome , senha from users where nome ='$nome' and senha = '$senha'";

            if($conn->query($SEARCH_SQL) === true){
                
                $_SESSION['nome'] = $nome;
                $_SESSION['senha'] = $senha;
                header("location:home.php");
            }
            else{
                echo "Nome ou senha incorretos";
            }
        }
    ?>
</body>
</html>