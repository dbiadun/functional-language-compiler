-- LED1 should blink every second, and LED2 should be switched on when the user button is pressed and switched off when it is not

main :: IO ();
main = sampleButtonInterrupt;

sampleButtonInterrupt :: IO ();
sampleButtonInterrupt = do {
    led <- tinyLED;
    led2 <- tinyLED2;
    button <- tinyBUTTON;

    outputMode <- tinyPinOutput;
    inputMode <- tinyPinInputPullup;

    tinyConfigure led outputMode;
    tinyConfigure led2 outputMode;
    tinyConfigure button inputMode;

    toggleChange <- tinyPinToggle;

    tinySetTimer 0 500000 (changeLed led);
    tinySetInterrupt button toggleChange (checkButtonCallback button led2);

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

checkButtonCallback :: Int -> Int -> Int -> IO ();
checkButtonCallback button led pin = checkButton button led;

checkButton :: Int -> Int -> IO ();
checkButton button led = do {
    -- We assume pullup button configuration
    buttonNotPressed <- tinyGet button;
    case not buttonNotPressed of {
        True -> tinyHigh led;
        False -> tinyLow led;
    };
};

not :: Bool -> Bool;
not b = case b of {
    True -> False;
    False -> True;
};

loop :: IO a -> IO a;
loop action = do {
    action;
    loop action;
};
