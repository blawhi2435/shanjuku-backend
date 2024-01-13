-- 新增欄位
ALTER TABLE public.groups
ADD COLUMN creator_id bigint NOT NULL;

-- 新增外鍵約束
ALTER TABLE public.groups
ADD CONSTRAINT fk_groups_user
FOREIGN KEY (creator_id)
REFERENCES public.users (id)
ON UPDATE NO ACTION
ON DELETE NO ACTION;