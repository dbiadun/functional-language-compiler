main :: IO ();
main = do {
    x <- getState;
    putInt x;
    putStr "\n";
    setState 15;
    x <- getState;
    putInt x;
    putStr "\n";
};

setState :: Int -> IO ();
setState val = stateSetInt "storedVal" val;

getState :: IO Int;
getState = stateGetInt "storedVal";