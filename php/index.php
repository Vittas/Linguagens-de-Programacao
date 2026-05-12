<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <form method="post">
        <select name="unit">
            <option value="f">Fahrenheit</option>
            <option value="k">Kelvin</option>
            <option value="c">Celsius</option>
        </select>
        <input type="number" name="temperature" placeholder="Digite a temperatura">
        <input type="submit" value="Converter">
    </form>

    <?php 
        if ($_SERVER["REQUEST_METHOD"] == "POST") {
            $temperature = $_POST["temperature"];
            $unit = $_POST["unit"];

            if ($unit == "f") {
                $celsius = ($temperature - 32) * 5/9;
                echo "$temperature °F é igual a $celsius °C";
            } elseif ($unit == "k") {
                $celsius = $temperature - 273.15;
                echo "$temperature K é igual a $celsius °C";
            } elseif ($unit == "c") {
                $fahrenheit = ($temperature * 9/5) + 32;
                $kelvin = $temperature + 273.15;
                echo "$temperature °C é igual a $fahrenheit °F e $kelvin K";
            }
        }
    ?>
</body>
</html>