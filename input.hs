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

putIntLine :: Int -> IO ();
putIntLine n = do {
    putInt n;
    putStr "\n";
};

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

main :: IO ();
main = sampleInfiniteFor;