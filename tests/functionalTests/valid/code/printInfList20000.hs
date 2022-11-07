import "List.hs";

main :: IO ();
main = do {
    printList numbers;
};

printList :: List Int -> IO ();
printList l = case l of {
    Nil -> putStr "";
    Cons x xs -> do {
        putInt x;
        case x >= 20000 of {
            False -> do {
                putStr ", ";
                printList xs;
            };
            True -> putStr "";
        };
    };
};

numbersFrom :: Int -> List Int;
numbersFrom n = Cons n (numbersFrom (n + 1));

numbers :: List Int;
numbers = numbersFrom 0;