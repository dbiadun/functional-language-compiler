main :: IO ();
main = do {
    putStr "Start\n";
    x <- return 5;
    y <- do {
        putStr "do\n";
        return 8;
    };
    putInt x;
    putStr "\n";
    putInt y;
    putStr "\n";
    z <- do {
        putStr "do2\n";
        return "do string";
    };
    putStr z;
    putStr "\n";
};