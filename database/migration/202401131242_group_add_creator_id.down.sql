-- 刪除外鍵約束
ALTER TABLE public.groups
DROP CONSTRAINT IF EXISTS fk_groups_user;

-- 刪除欄位
ALTER TABLE public.groups
DROP COLUMN IF EXISTS creator_id;