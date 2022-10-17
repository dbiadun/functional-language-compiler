main :: IO ();
main = do {
    x <- fun;
    putStr x;
};

fun :: IO ();
fun = do {
    putStr "fun\n";
};
