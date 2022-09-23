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

blink :: Int -> Int -> IO ();
blink led time = do {
    halfTime <- return (time / 2);
    singleSleep <- return (halfTime / 10);

    -- We use many short sleeps to enable good interruption handling
    tinyHigh led;
    for 0 singleSleep (takeInt (tinySleep 10));
    tinyLow led;
    for 0 singleSleep (takeInt (tinySleep 10));
};

not :: Bool -> Bool;
not b = case b of {
    True -> False;
    False -> True;
};

checkButton :: Int -> Int -> IO ();
checkButton button led = do {
    -- We assume pullup button configuration
    putStr "INTERRUPTION\n\r";
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
    loop (blink led 1000);
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

    tinySetInterrupt button toggleChange (checkButtonCallback button led);

    --loop (putStr "");
    loop (blink led2 500);
};

main :: IO ();
main = sampleInterruptionButton;