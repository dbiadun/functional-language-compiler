main :: IO ();
main = do {
    printBool True;
    printBool False;
};

printBool :: Bool -> IO ();
printBool b = case b of {
    True -> putStr "True\n";
    x -> putStr "False\n";
};