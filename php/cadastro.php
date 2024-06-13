<?php
    include 'conexaobanco.php'
?>

<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="CSS/cadastro.css">
    <title>Document</title>
</head>
<body>
    <nav>Cadastro</nav>
    <br>
    <div>
        <form method="post">
            <h1>Cadastre-se</h1>
            <div id="cadform">
                <label for="cad_nome">Nome de usuário</label>
                <input type="text" name="nome" placeholder="Digite seu nome" required>
                <label for="cad_senha">Senha</label>
                <input type="password" name="senha" id="" placeholder="Digite sua senha" required>
                <br>
                <input type="submit" id="bt_enviar" name="submit" value="submit">
            </div>
        </form>
    </div>

    <?php
            
            if(isset($_POST['submit'])){
                $nome =  $_POST['nome'];
                $senha = $_POST['senha'];
                $status = TRUE;
                


                $SEND_SQL = "INSERT INTO users VALUES (null,'$nome','$senha', '$status')";
            
                if($conn->query($SEND_SQL) === true){

                    $_SESSION['nome'] = $nome;
                    $_SESSION['senha'] = $senha;
                    header("location:home.php");
                }
                else{
                    echo "erro";
                }
            }
    ?>
    
</body>
</html>
