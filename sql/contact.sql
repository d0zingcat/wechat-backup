DROP TABLE IF EXISTS contact;
CREATE TABLE contact (
    username VARCHAR(255) PRIMARY KEY,
    nickname VARCHAR(255), 
    alias VARCHAR(255),
    username_pinyin VARCHAR(255),
    username_pinyin_short VARCHAR(128),
    remark VARCHAR(255),
    remark_pinyin VARCHAR(255),
    remark_pinyin_short VARCHAR(128),
    sex SMALLINT,
    head_image_url TEXT,
    head_image_url_hd TEXT,
    brand_icon_url TEXT
);
