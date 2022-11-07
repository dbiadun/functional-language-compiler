import "Maybe.hs";
import "List.hs";
-- Line comment

{- Multiline comment
-- Still a comment
-}

id :: b -> b;
id x = x;

append :: a -> List a -> List a;
append x = (id appendToList) x;

add :: Int -> Int -> Int;
add a b = a + b;

sum :: List Int -> Int;
sum = foldl add 0;

just :: Maybe Int -> Int;
just x = case x of {
    Just n -> n;
    Nothing -> 0;
};

sumOfFirsts :: List (List Int) -> Int;
sumOfFirsts l = sum (map just (listOfFirsts l));

boolToInt :: Bool -> Int;
boolToInt b = case b of {
    False -> 0;
    True -> 1;
};

numbersFrom :: Int -> List Int;
numbersFrom n = Cons n (numbersFrom (n + 1));

numbers :: List Int;
numbers = numbersFrom 0;

nNumbers :: Int -> List Int;
nNumbers n = take n numbers;

putBool :: Bool -> IO ();
putBool b = case b of {
    True -> putStr "True";
    False -> putStr "False";
};

printIntList :: List Int -> IO ();
printIntList l = do {
    putStr "[";
    printIntListHelper l;
    putStr "]";
};

printIntListHelper :: List Int -> IO ();
printIntListHelper l = case l of {
   Nil -> putStr "";
   Cons x xs -> case xs of {
       Nil -> putInt x;
       l -> do {
           putInt x;
           putStr ", ";
           printIntListHelper l;
       };
   };
};

getName :: IO String;
getName = do {
    putStr "What's your name?\n";
    name <- getLine;
    putStr "Thank you!\n";
    return name;
};

loop :: IO a -> IO a;
loop action = do {
    action;
    loop action;
};

printHello :: IO ();
printHello = putStr "Hello\n";

printNumbersFrom :: Int -> IO ();
printNumbersFrom n = do {
    putInt n;
    putStr "\n";
    printNumbersFrom (n + 1);
};

forInf :: Int -> (Int -> IO ()) -> IO ();
forInf n fun = do {
    fun n;
    forInf (n + 1) fun;
};

for :: Int -> Int -> (Int -> IO ()) -> IO ();
for i n fun = case i < n of {
    False -> putStr "";
    True -> do {
        fun i;
        for (i + 1) n fun
    };
};

takeInt :: IO () -> (Int -> IO ());
takeInt action n = action;

putIntLine :: Int -> IO ();
putIntLine n = do {
    putInt n;
    putStr "\n";
};

blinkLoop :: Int -> Int -> IO ();
blinkLoop led time = do {
    halfTime <- return (time / 2);
    tinySetTimer 0 halfTime (changeLed led);
};

changeLed :: Int -> IO ();
changeLed led = do {
    ledState <- tinyGet led;

    case ledState of {
        True -> tinyLow led;
        False -> tinyHigh led;
    };
};

not :: Bool -> Bool;
not b = case b of {
    True -> False;
    False -> True;
};

checkButton :: Int -> Int -> IO ();
checkButton button led = do {
    -- We assume pullup button configuration
    buttonNotPressed <- tinyGet button;
    case not buttonNotPressed of {
        True -> tinyHigh led;
        False -> tinyLow led;
    };
};

checkButtonCallback :: Int -> Int -> Int -> IO ();
checkButtonCallback button led pin = checkButton button led;

sampleSumOfFirsts :: Int;
sampleSumOfFirsts = sumOfFirsts (Cons (Cons 4 (Cons 7 Nil)) (Cons Nil (Cons (Cons 2 (Cons 3 Nil)) Nil)));

sampleLazyOperators :: Int;
sampleLazyOperators = just (fst (Cons 3 (Cons (1/0) Nil)));

sampleTake :: Int; -- Checks lazy evaluation of infinite structures like `numbers`
sampleTake = sum (take 4 (Cons 3 (Cons 2 (Cons 4 (Cons 6 (Cons 7 Nil)))))); -- 15


listOfFirsts :: List (List a) -> List (Maybe a);
listOfFirsts = map fst;

listOfLasts :: List (List a) -> List (Maybe a);
listOfLasts = map last;

sampleSumOfLasts :: Int; -- Checks the performance when many nodes are needed (garbage collection)
sampleSumOfLasts = sum (map just (listOfLasts (map nNumbers (nNumbers 30))));

samplePrintList :: IO ();
samplePrintList = do {
    printIntList (nNumbers 20);
    putStr "\n";
};

samplePrintInfList :: IO (); -- Should not stop
samplePrintInfList = do {
    printIntList numbers;
    putStr "\n";
};

sampleGreeting :: IO ();
sampleGreeting = do {
    name <- getName;
    putStr "Hi, ";
    putStr name;
    putStr "!\n";
};

samplePrintNumbers :: IO ();
samplePrintNumbers = printNumbersFrom 0;

sampleInfiniteFor :: IO ();
sampleInfiniteFor = forInf 0 putIntLine;

sampleBlinkingLed :: IO ();
sampleBlinkingLed = do {
    led <- tinyLED;
    outputMode <- tinyPinOutput;
    tinyConfigure led outputMode;
    blinkLoop led 1000000;
};

sampleSimpleButton :: IO ();
sampleSimpleButton = do {
    led <- tinyLED;
    led2 <- tinyLED2;
    button <- tinyBUTTON;

    outputMode <- tinyPinOutput;
    inputMode <- tinyPinInputPullup;

    tinyConfigure led outputMode;
    tinyConfigure led2 outputMode;
    tinyConfigure button inputMode;

    loop (checkButton button led2);
};

sampleInterruptionButton :: IO ();
sampleInterruptionButton = do {
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

timerStopStart :: Int -> IO ();
timerStopStart timer = do {
    tinyStopTimer timer;
    tinySleep 3000;
    tinyStartTimer timer;
};

sampleStopStartTimer :: IO ();
sampleStopStartTimer = do {
    led <- tinyLED;
    led2 <- tinyLED2;
    button <- tinyBUTTON;

    outputMode <- tinyPinOutput;
    inputMode <- tinyPinInputPullup;

    tinyConfigure led outputMode;
    tinyConfigure led2 outputMode;
    tinyConfigure button inputMode;

    buttonPressChange <- tinyPinFalling;

    timerNum <- return 0;

    tinySetTimer timerNum 500000 (changeLed led);
    tinySetInterrupt button buttonPressChange (takeInt (timerStopStart timerNum));

    tinyHigh led2;
    loop (putStr "");

};

runTimerCallbackUntil :: Int -> Int -> IO () -> IO ();
runTimerCallbackUntil timer ticks action = do {
    curTick <- stateGetInt "curTick";

    case curTick < ticks of {
        True -> do {
            stateSetInt "curTick" (curTick + 1);
            action;
        };
        False -> tinyStopTimer timer;
    };
};

timerSetNTicks :: Int -> Int -> Int -> IO () -> IO ();
timerSetNTicks timer duration ticks action = do {
    stateSetInt "curTick" 0;
    tinySetTimer timer duration (runTimerCallbackUntil timer ticks action);
};

sampleState :: IO ();
sampleState = do {
    tinySleep 2000;
    led <- tinyLED;
    led2 <- tinyLED2;

    outputMode <- tinyPinOutput;

    tinyConfigure led outputMode;
    tinyConfigure led2 outputMode;

    timerSetNTicks 0 500000 (2 * 5) (changeLed led);

    tinyHigh led2;
    loop (putStr "");
};

getNth :: Int -> List a -> a;
getNth n l = case n > 0 of {
    True -> case l of {
        Cons x xs -> getNth (n - 1) xs;
    };
    False -> case l of {
        Cons x xs -> x;
    };
};

readTemp :: Int -> Int -> IO ();
readTemp i2c termometer = do {
    res <- tinyI2CTx i2c termometer Nil 6;
    case res of {
        Left err -> putStr "";
        Right l -> do {
            tinyStopTimer 0;

            tempH <- return (getNth 0 l);
            tempL <- return (getNth 1 l);
            humH <- return (getNth 3 l);
            humL <- return (getNth 4 l);

            tempOutput <- return (256 * tempH + tempL);
            humOutput <- return (256 * humH + humL);

            temp <- return (10 * (0 - 45) + 10 * 175 * tempOutput / (256 * 256 - 1));
            tempInt <- return (temp / 10);
            tempRest <- return (temp - 10 * tempInt);
            hum <- return (100 * humOutput / (256 * 256 - 1));

            putStr "Temperature: ";
            putInt tempInt;
            putStr ",";
            putInt tempRest;
            putStr "\n";
            putStr "Humidity: ";
            putInt hum;
            putStr "%\n";
        };
    };
};

sampleTermometer :: IO ();
sampleTermometer = do {
    tinySleep 2000;

    termometer <- return 0x44;
    i2c <- return 0;

    tinyI2CConfigureDefault i2c;

    w <- return (Cons 0x24 (Cons 0x00 Nil));
    tinyI2CTx i2c termometer w 0;

    tinySetTimer 0 100000 (readTemp i2c termometer);

    loop (putStr "");
};

sampleBitOps :: IO ();
sampleBitOps = do {
    tinySleep 2000;
    putInt (3 | 5 | 6); -- 7
    putStr "\n";
    putInt (7 & 14 & 22); -- 6
    putStr "\n";
    putInt (shiftL 6 2); -- 24
    putStr "\n";
    putInt (shiftR 29 3); -- 3
    putStr "\n";
};

right :: Either a b -> b;
right x = case x of {
    Right x -> x;
};

twoBytesToInt :: List Int -> Int;
twoBytesToInt l = 256 * (getNth 1 l) + (getNth 0 l);

readTempSecond :: Int -> Int -> IO ();
readTempSecond i2c termometer = do {
    statusReg <- return 0xF3; -- 8 bits
    statusMeasuringBit <- return (shiftL 1 3);

    statusEither <- tinyI2CReadRegister i2c termometer statusReg 1;
    status <- return (right statusEither);
    measured <- return (((getNth 0 status) & statusMeasuringBit) == 0);

    case measured of {
        False -> putStr "";
        True -> do {
            tinyStopTimer 0;

            tempStartReg <- return 0xFA; -- 20 bits
            t1StartReg <- return 0x88; -- 16 bits
            t2StartReg <- return 0x8A; -- 16 bits
            t3StartReg <- return 0x8C; -- 16 bits

            tempEither <- tinyI2CReadRegister i2c termometer tempStartReg 3;
            tempBytes <- return (right tempEither);
            tempMsb <- return (getNth 0 tempBytes);
            tempLsb <- return (getNth 1 tempBytes);
            tempXlsb <- return (shiftR (getNth 2 tempBytes) 4);
            tempOutput <- return (256 * 16 * tempMsb + 16 * tempLsb + tempXlsb);

            t1Either <- tinyI2CReadRegister i2c termometer t1StartReg 2;
            digT1 <- return (twoBytesToInt (right t1Either));

            t2Either <- tinyI2CReadRegister i2c termometer t2StartReg 2;
            digT2 <- return (twoBytesToInt (right t2Either));
            digT2 <- case digT2 >= 256 * 128 of {
                False -> return digT2;
                True -> return (digT2 - 256 * 256);
            };

            t3Either <- tinyI2CReadRegister i2c termometer t3StartReg 2;
            digT3 <- return (twoBytesToInt (right t3Either));
            digT3 <- case digT3 >= 256 * 128 of {
                False -> return digT3;
                True -> return (digT3 - 256 * 256);
            };

            var1 <- return ((1000 * tempOutput / 16384 - 1000 * digT1 / 1024 ) * digT2 / 1000);
            varMid <- return (100 * tempOutput / 131072 - 100 * digT1 / 8192);
            var2 <- return (varMid * varMid * digT3 / (100 * 100));
            tFine <- return (var1 + var2);
            t <- return (tFine / 512);
            tempInt <- return (t / 10);
            tempRest <- return (t - 10 * tempInt);

            putStr "Temperature: ";
            putInt tempInt;
            putStr ",";
            putInt tempRest;
            putStr "\n";
        };
    };
};

sampleSecondTermometer :: IO ();
sampleSecondTermometer = do {
    tinySleep 2000;

    i2c <- return 0;
    termometer <- return 0x77;

    tinyI2CConfigureDefault i2c;

    ctrlReg <- return 0xF4; -- 8 bits

    osrsT <- return (shiftL 1 5);
    mode <- return 1;
    ctrlRegVal <- return (osrsT | mode);


    tinyI2CWriteRegister i2c termometer ctrlReg (Cons ctrlRegVal Nil);

    tinySetTimer 0 500000 (readTempSecond i2c termometer);

    loop (putStr "");
};

main :: IO ();
main = sampleSecondTermometer;
