import "List.hs";
import "../../libs/SortType.hs";

data ListPair = ListPair (List SortType) (List SortType);

main :: IO ();
main = do {
    tinySleep 2000;

    l <- return (getList 16);
    l <- return (mergeSort l);

    printList l;
    putStr "\n";
};

mergeSort :: List SortType -> List SortType;
mergeSort l = case l of {
    Nil -> Nil;
    Cons x xs -> case xs of {
        Nil -> l;
        Cons y ys -> case split l (ListPair Nil Nil) of {
            ListPair left right -> merge (mergeSort left) (mergeSort right);
        };
    };
};

split :: List SortType -> ListPair -> ListPair;
split l pair = case pair of {
    ListPair left right -> case l of {
        Nil -> pair;
        Cons x xs -> case xs of {
            Nil -> ListPair (Cons x left) right;
            Cons y ys -> split ys (ListPair (Cons x left) (Cons y right));
        };
    };
};

merge :: List SortType -> List SortType -> List SortType;
merge left right = case left of {
    Nil -> right;
    Cons x xs -> case right of {
        Nil -> left;
        Cons y ys -> case cmp x y of {
            True -> Cons x (merge xs right);
            False -> Cons y (merge left ys);
        };
    };
};
