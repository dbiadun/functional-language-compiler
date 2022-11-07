import "List.hs";

main :: IO ();
main = do {
    numbers <- return (numbersFrom 0);
    printIntList (take 0 numbers);
    putStr "\n";
    printIntList (take 1 numbers);
    putStr "\n";
    printIntList (take 3 numbers);
    putStr "\n";
    printIntList (take 6 numbers);
    putStr "\n";
};

numbersFrom :: Int -> List Int;
numbersFrom n = Cons n (numbersFrom (n + 1));

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
        list -> do {
            putInt x;
            putStr ", ";
            printIntListHelper xs;
        };
    };
};
