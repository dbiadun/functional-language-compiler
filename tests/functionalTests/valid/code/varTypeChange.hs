main :: IO ();
main = do {
    x <- return 5;
    putInt x;
    putStr "\n";
    x <- return "string var";
    putStr x;
    putStr "\n";
};