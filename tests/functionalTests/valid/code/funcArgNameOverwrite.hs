main :: IO ();
main = do {
    arg <- return 5;
    putInt arg;
    putStr "\n";
    fun arg;
    putInt arg;
    putStr "\n";
};

fun :: Int -> IO ();
fun arg = do {
    arg <- return 13;
    putInt arg;
    putStr "\n";
};