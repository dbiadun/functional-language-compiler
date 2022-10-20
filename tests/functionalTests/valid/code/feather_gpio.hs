main :: IO ();
main = do {
    tinySleep 2000;

    led1 <- tinyLED;
    led2 <- tinyLED2;

    outputMode <- tinyPinOutput;

    tinyConfigure led1 outputMode;
    tinyConfigure led2 outputMode;

    printLED led1;
    printLED led2;
    putStr "\n";

    tinyHigh led1;
    printLED led1;
    printLED led2;
    putStr "\n";

    tinyHigh led2;
    printLED led1;
    printLED led2;
    putStr "\n";

    tinyLow led1;
    printLED led1;
    printLED led2;
    putStr "\n";

    tinyLow led2;
    printLED led1;
    printLED led2;
};

printLED :: Int -> IO ();
printLED led = do {
    on <- tinyGet led;
    case on of {
        True -> putStr "ON\n";
        False -> putStr "OFF\n";
    };
};
