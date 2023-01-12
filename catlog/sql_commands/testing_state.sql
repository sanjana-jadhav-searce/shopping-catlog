DELETE FROM cart_item;
DELETE FROM product;
DELETE FROM cart_reference;
DELETE FROM category;

INSERT INTO category VALUES(1, 'Medical');
INSERT INTO category VALUES(2, 'Electronics');
INSERT INTO category VALUES(3, 'Accessories');

INSERT INTO product VALUES(1, 'Vatika', '{"color": "green", "gender": "female"}', 67283, 1, 54.56);
INSERT INTO product VALUES(2, 'Jacket', '{"color": "blue", "H&M": "true"}', 37348, 2, 23.38);
INSERT INTO product VALUES(3, 'Jeans', '{"color": "Black"}', 37472, 1, 30.99);

INSERT INTO inventory VALUES(3, 500);

INSERT INTO cart_item VALUES('23d783d-7663-451d-b79e-dsjnd73szn', 3, 10);

INSERT INTO cart_reference VALUES('23d783d-7663-451d-b79e-dsjnd73szn', '2022-12-28 15:52:22.013532');