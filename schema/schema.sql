USE `auster`;

CREATE TABLE `user`
(
    `id`           varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`         varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `gender`       varchar(10) COLLATE utf8mb4_bin  NOT NULL,
    `age`          int(11)                          NOT NULL,
    `profile_path` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='ユーザ';

CREATE TABLE `hobby`
(
    `id`   varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='趣味';

INSERT INTO auster.hobby (id, name) VALUES ('cstkdiat6c3011a83so0', '釣り');
INSERT INTO auster.hobby (id, name) VALUES ('cstkdiat6c3011a83sog', 'キャンプ');

CREATE TABLE `user_hobby`
(
    `user_id`  varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `hobby_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`user_id`, `hobby_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='ユーザの趣味';

CREATE TABLE `vendor`
(
    `id`      varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行会社';

INSERT INTO auster.vendor (id, name, address) VALUES ('can1', 'CampsiteTORAMI', '千葉県長生郡一宮町東浪見1611');
INSERT INTO auster.vendor (id, name, address) VALUES ('can2', 'Beach Camp 九十九里', '千葉県大網白里市四天木2761-40');
INSERT INTO auster.vendor (id, name, address) VALUES ('can3', '銚子電気鉄道', '千葉県銚子市新生町2丁目297番地');
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo1', '銚子市漁業協同組合', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo2', '銚子市生活環境課 清掃美化班', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）');
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo3', '銚子市観光課', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）');

CREATE TABLE `travel_spot`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `vendor_id`   varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `title`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `take_time`   int(11)                          NOT NULL COMMENT '所要時間（分）',
    `description` text COLLATE utf8mb4_bin        NOT NULL COMMENT '説明',
    `address`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `photo_path`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `level`       int NOT NULL COMMENT '難易度',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポット';

-- 体験スポット
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('1', 'gyo1', '釣り体験 ヒラマサ', 180, '船釣りでヒラマサを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/1/1.jpg', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('10', 'can2', 'ブッシュクラフト火おこし体験', 300, 'キャンパーが一度は憧れるブッシュクラフト、火おこし体験ができます', '千葉県大網白里市四天木2761-40', '/assets/images/travel_spots/10/10.jpg', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('11', 'can1', 'トレインランニング', 300, '銚子電鉄の駅を自分の足で巡りながら大自然を駆け抜けるトレインランニングを体験できます', '千葉県長生郡一宮町東浪見1611', '/assets/images/travel_spots/11/11.jpg', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('12', 'can3', 'キャンプ場のイベントで1日バーテンダー', 180, '一日バーテンダーをやって、キャンプ場を盛り上げる体験ができます', '千葉県銚子市新生町2丁目297番地', '/assets/images/travel_spots/12/12.png', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('13', 'can2', 'ショップの企画担当', 30, 'キャンプ場周辺のショップで企画担当を募集します', '千葉県大網白里市四天木2761-40', '/assets/images/travel_spots/13/13.jpg', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('14', 'can3', 'スペシャルオファー<複業> 自然と音楽の交差点', 120, '地域の人とキャンパーがそこでしか作れない音楽体験を提供してみませんか？', '千葉県大網白里市四天木2761-40', '/assets/images/travel_spots/14/14.jpg', 0);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('2', 'gyo1', '釣り体験 シイラ', 240, '船釣りでシイラを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/2/2.jpg', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('3', 'gyo2', '社会貢献活動 浜辺のゴミ拾い', 90, '銚子の海岸に打ち上げられた漂流物の清掃を実施します', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）', '/assets/images/travel_spots/3/3.png', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('4', 'gyo1', '漁業体験 定置網漁', 180, '定置網漁を体験できます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/4/4.png', 1);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('5', 'gyo1', '釣り体験 タイ', 180, '弾きが強いタイを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/5/5.jpg', 2);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('6', 'gyo1', '釣り体験 太刀魚', 180, '船釣りで太刀魚を釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/6/6.jpg', 2);
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path, level) VALUES ('7', 'gyo2', 'スペシャルオファー<複業> 漁師の思いを発信', 120, '銚子の漁業組合でマーケターとして働き、漁師さんの誇りと情熱をより多くの人に伝えてみませんか？', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）', '/assets/images/travel_spots/7/7.png', 0);

CREATE TABLE `travel_spot_itinerary`
(
    `id`             varchar(20) COLLATE utf8mb4_bin   NOT NULL,
    `travel_spot_id` varchar(20) COLLATE utf8mb4_bin   NOT NULL,
    `title`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `description`    text COLLATE utf8mb4_bin         NOT NULL,
    `kind`           enum ('experience','spot','move') NOT NULL COMMENT 'experience: 体験場所, spot: 観光地, move: 移動',
    `take_time`      int(11)                           NOT NULL COMMENT '所要時間（分）',
    `price`          int(11)                           NOT NULL,
    `order`          int(11)                           NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポットと旅程の関連';

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti1','特急 しおさい３号','', 'move', '1', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti2','犬吠埼灯台', '銚子のシンボルとして知られる白亜の灯台です。本州で最も早く朝日が昇る場所としても有名です。(千葉県銚子市犬吠埼957612)','spot', '1', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti3','徒歩', '','move', '1', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti4','釣り体験 ヒラマサ', '','experience', '1', 180, 15000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti5','特急 しおさい14号', '','move', '1', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti6','特急 しおさい３号','', 'move', '2', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti7','地球の丸く見える丘展望館', '標高90mの屋上展望スペースからは360度の大パノラマが楽しめます。水平線の両端が丸みを帯びて見えることから、その名が付けられました。(千葉県銚子市天王台1421-13)','spot', '2', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti8','徒歩', '','move', '2', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti9','釣り体験 シイラ', '','experience', '2', 180, 15000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti10','特急 しおさい14号', '','move', '2', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti11','特急 しおさい３号','', 'move', '23', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti12','圓福寺（飯沼観音）', '真言宗の寺院で、貴重な美術品や古文書が保管されています。特に「釈迦涅槃殿」は千葉県の重要文化財に指定されています。(千葉県銚子市馬場町2934)','spot', '3', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti13','徒歩', '','move', '3', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti14','釣り体験 ヒラメ','', 'experience', '3', 180, 15000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti15','特急 しおさい14号', '','move', '3', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti16','特急 しおさい３号','', 'move', '4', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti17','犬吠埼灯台', '銚子のシンボルとして知られる白亜の灯台です。本州で最も早く朝日が昇る場所としても有名です。(千葉県銚子市犬吠埼957612)','spot', '4', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti18','徒歩', '','move', '4', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti19','漁業体験 競りの見学','', 'experience', '4', 120, 5000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti20','特急 しおさい14号', '','move', '4', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti21','特急 しおさい３号','', 'move', '5', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti22','地球の丸く見える丘展望館', '標高90mの屋上展望スペースからは360度の大パノラマが楽しめます。水平線の両端が丸みを帯びて見えることから、その名が付けられました。(千葉県銚子市天王台1421-13)','spot', '5', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti23','徒歩', '','move', '5', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti24','漁業体験 定置網漁','', 'experience', '5', 180, 12000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti25','特急 しおさい14号', '','move', '5', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti26','特急 しおさい３号','', 'move', '6', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti27','圓福寺（飯沼観音）', '真言宗の寺院で、貴重な美術品や古文書が保管されています。特に「釈迦涅槃殿」は千葉県の重要文化財に指定されています。(千葉県銚子市馬場町2934)','spot', '6', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti28','徒歩', '','move', '6', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti29','社会貢献活動 浜辺のゴミ拾い','', 'experience', '6', 120, 0, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti30','特急 しおさい14号', '','move', '6', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti31','特急 しおさい３号','', 'move', '10', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti32','犬吠埼灯台', '銚子のシンボルとして知られる白亜の灯台です。本州で最も早く朝日が昇る場所としても有名です。(千葉県銚子市犬吠埼957612)','spot', '10', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti33','徒歩', '','move', '10', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti34','キャンプ場運営体験 イベントのスタッフ','', 'experience', '10', 300, 8000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti35','特急 しおさい14号', '','move', '10', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti36','特急 しおさい３号','', 'move', '11', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti37','地球の丸く見える丘展望館', '標高90mの屋上展望スペースからは360度の大パノラマが楽しめます。水平線の両端が丸みを帯びて見えることから、その名が付けられました。(千葉県銚子市天王台1421-13)','spot', '11', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti38','徒歩', '','move', '11', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti39','キャンプ中級者体験 ブッシュクラフト火おこし','', 'experience', '11', 60, 5000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti40','特急 しおさい14号', '','move', '11', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti41','特急 しおさい３号','', 'move', '12', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti42','圓福寺（飯沼観音）', '真言宗の寺院で、貴重な美術品や古文書が保管されています。特に「釈迦涅槃殿」は千葉県の重要文化財に指定されています。(千葉県銚子市馬場町2934)','spot', '12', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti43','徒歩', '','move', '12', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti44','トレインランニング','', 'experience', '12', 300, 3000, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti45','特急 しおさい14号', '','move', '12', 60, 1500, 5);

INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti46','特急 しおさい３号','', 'move', '13', 60, 1500, 1);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti47','犬吠埼灯台', '銚子のシンボルとして知られる白亜の灯台です。本州で最も早く朝日が昇る場所としても有名です。(千葉県銚子市犬吠埼957612)','spot', '13', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti48','徒歩', '','move', '13', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti49','スペシャルオファー<複業> ショップの企画担当','', 'experience', '13', 0, 0, 4);
INSERT INTO auster.travel_spot_itinerary (id, title, description, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti50','特急 しおさい14号', '','move', '13', 60, 1500, 5);


CREATE TABLE `travel_spot_itinerary_item`
(
    `id`                       varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `travel_spot_itinerary_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`                     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `number`                   int(11)                          NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポットで必要な持ち物';

-- 釣り体験（ヒラマサ）の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_1', 'iti4', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_2', 'iti4', '日焼け止め', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_3', 'iti4', 'タオル', 2);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_4', 'iti4', '酔い止め薬', 1);

-- 釣り体験（シイラ）の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_1', 'iti9', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_2', 'iti9', '日焼け止め', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_3', 'iti9', 'タオル', 2);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_4', 'iti9', '酔い止め薬', 1);

-- 釣り体験（ヒラメ）の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_1', 'iti14', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_2', 'iti14', '日焼け止め', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_3', 'iti14', 'タオル', 2);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_4', 'iti14', '酔い止め薬', 1);

-- 競りの見学の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item4_1', 'iti19', '長靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item4_2', 'iti19', '防寒着', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item4_3', 'iti19', 'メモ帳', 1);

-- 定置網漁体験の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_1', 'iti24', '長靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_2', 'iti24', '作業用手袋', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_3', 'iti24', '防寒着', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_4', 'iti24', 'タオル', 2);

-- 浜辺のゴミ拾いの持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item6_1', 'iti29', '軍手', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item6_2', 'iti29', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item6_3', 'iti29', '飲み物', 1);

-- キャンプ場運営体験の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item10_1', 'iti34', '作業用手袋', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item10_2', 'iti34', '動きやすい服装', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item10_3', 'iti34', '長靴', 1);

-- キャンプ中級者体験の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item11_1', 'iti39', '軍手', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item11_2', 'iti39', '防虫スプレー', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item11_3', 'iti39', '長袖の服', 1);

-- トレインランニングの持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item12_1', 'iti44', 'ランニングシューズ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item12_2', 'iti44', '水筒', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item12_3', 'iti44', 'タオル', 2);

-- 複業オファーの持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item13_1', 'iti49', '筆記用具', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item13_2', 'iti49', '履歴書', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item13_3', 'iti49', 'ポートフォリオ', 1);

-- 犬吠埼灯台見学の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_1', 'iti2', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_2', 'iti2', '日傘・帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_3', 'iti2', '歩きやすい靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_4', 'iti17', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_5', 'iti17', '日傘・帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_6', 'iti17', '歩きやすい靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_7', 'iti32', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_8', 'iti32', '日傘・帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_9', 'iti32', '歩きやすい靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_10', 'iti47', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_11', 'iti47', '日傘・帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_12', 'iti47', '歩きやすい靴', 1);

-- 地球の丸く見える丘展望館の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_1', 'iti7', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_2', 'iti7', '双眼鏡', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_3', 'iti22', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_4', 'iti22', '双眼鏡', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_5', 'iti37', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_6', 'iti37', '双眼鏡', 1);

-- 圓福寺見学の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_1', 'iti12', '歩きやすい靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_2', 'iti12', '写真撮影許可が必要なカメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_3', 'iti12', 'お賽銭', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_4', 'iti27', '歩きやすい靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_5', 'iti27', '写真撮影許可が必要なカメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_6', 'iti27', 'お賽銭', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_7', 'iti42', '歩きやすい靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_8', 'iti42', '写真撮影許可が必要なカメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_9', 'iti42', 'お賽銭', 1);

CREATE TABLE `travel_spot_diary`
(
    `id`             varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `user_id`        varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `travel_spot_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `title`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `date`           date                             NOT NULL,
    `photo_path`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `description`    text COLLATE utf8mb4_bin         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='日記';

CREATE TABLE `travel_spot_hobby`
(
    `travel_spot_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `hobby_id`       varchar(20) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`travel_spot_id`, `hobby_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='体験スポットに当てはまる趣味';

INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('1', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('2', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('3', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('4', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('5', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('6', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('10', 'cstkdiat6c3011a83sog');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('11', 'cstkdiat6c3011a83sog');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('12', 'cstkdiat6c3011a83sog');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('13', 'cstkdiat6c3011a83sog');

CREATE TABLE `reservation`
(
    `id`                   varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `user_id`              varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `travel_spot_id`       varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `travel_spot_diary_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `from_date`            date                            NOT NULL,
    `to_date`              date                            NOT NULL,
    `is_offer`             tinyint(1)                      NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='予約';

CREATE TABLE `encounter`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`        varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `place`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `user_id`     varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `date`        date                             NOT NULL,
    `description` text COLLATE utf8mb4_bin         NOT NULL,
    PRIMARY KEY (`id`),
    KEY           `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='出会った人';
