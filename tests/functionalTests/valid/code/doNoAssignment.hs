main :: IO ();
main = do {
    funcWithReturn;
    putStr "";
};

funcWithReturn :: IO Int;
funcWithReturn = do {
    putStr "funcWithReturn\n";
    return 247;
};