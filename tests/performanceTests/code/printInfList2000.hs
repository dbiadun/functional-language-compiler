import "List.hs";

main :: IO ();
main = do {
    tinySleep 2000;

    printList numbers;
};

printList :: List Int -> IO ();
printList l = case l of {
    Nil -> putStr "";
    Cons x xs -> do {
        putInt x;
        putStr ", ";
        printList xs;
    };
};

numbersFrom :: Int -> List Int;
numbersFrom n = Cons n (numbersFrom (n + 1));

numbers :: List Int;
numbers = numbersFrom 0;