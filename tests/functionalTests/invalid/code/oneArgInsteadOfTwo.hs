main :: IO ();
main = do {
    fun 5;
};

fun :: Int -> Int -> IO ();
fun x y = do {
    putInt x;
    putStr "\n";
};
