-- The LED on the board should blink every second

main :: IO ();
main = sampleBlinkingLed;

sampleBlinkingLed :: IO ();
sampleBlinkingLed = do {
    led <- tinyLED;
    outputMode <- tinyPinOutput;
    tinyConfigure led outputMode;
    blinkLoop led 1000000;

};

blinkLoop :: Int -> Int -> IO ();
blinkLoop led time = do {
    halfTime <- return (time / 2);
    tinySetTimer 0 halfTime (changeLed led);
    loop (putStr "");
};

changeLed :: Int -> IO ();
changeLed led = do {
    ledState <- tinyGet led;

    case ledState of {
        True -> tinyLow led;
        False -> tinyHigh led;
    };
};

loop :: IO a -> IO a;
loop action = do {
    action;
    loop action;
};
