main :: IO ();
main = do {
    setState 15;
    x <- getState;
    putInt x;
};

setState :: Int -> IO ();
setState val = stateSetInt "storedVal" val;

getState :: IO Int;
getState = stateGetInt "storedVal";