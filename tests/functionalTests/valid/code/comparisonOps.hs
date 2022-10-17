main :: IO ();
main = do {
    printBool (3 < 5);
    printBool (3 > 5);
    printBool (15 <= 16);
    printBool (16 <= 16);
    printBool (17 <= 16);
    printBool (10 >= 10);
    printBool (10 >= 11);
    printBool (13 == 13);
    printBool (13 == 14);
    printBool (15 /= 20);
    printBool (15 /= 15);
};

printBool :: Bool -> IO ();
printBool b = case b of {
    True -> putStr "True\n";
    False -> putStr "False\n";
};