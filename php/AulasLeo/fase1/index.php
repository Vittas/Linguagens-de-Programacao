<?php

    $times = [ "Time A", "Time B", "Time C", "Time D" ];

?>

<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    <table border="2">
        <tr>
            <?php 
                foreach(range(1,12) as $num) {
                    echo "<td>Rodada $num</td>";
                }
            ?>
        </tr>
        <tr>
            <?php
                foreach($times as $time) {
                    foreach($times as $time2){
                        if($time !== $time2) {
                            echo "<td>$time vs $time2</td>";
                        }
                    }
                }
            ?>
        </tr>
    </table>

</body>

</html>