main :: IO ();
main = do {
    putStr "Start\n";
    x <- do {
        putStr "do\n";
    };
    putInt x;
    putStr "End\n";
};