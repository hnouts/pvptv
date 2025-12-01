INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('21faf757-3065-ad90-2a5c-f2d24fd271a5', 'volimond', 'Volimond', true, '{"discord":"https://discord.gg/irontriangle","instagram":"https://www.instagram.com/volimond/","twitter":"https://twitter.com/Tank4daze"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '21faf757-3065-ad90-2a5c-f2d24fd271a5', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'vengeance' AND c.class_slug = 'demon_hunter'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('3fe9d8c8-01ec-ae16-ced1-f34c4ba96b31', 'fuseton', 'Fuseton', true, '{"twitter":"https://twitter.com/fuseton"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('81e7a6ca-077d-ca1f-b257-9d8888d29bfd', 'exzistancetw', 'Exzistancetw', true, '{"discord":"https://discord.com/invite/BvHMFpY","twitter":"https://twitter.com/whocareslul","youtube":"https://www.youtube.com/channel/UCQU6XlJgVKMDRv3NnqgdIJw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('3ce4a46e-0ea4-36c8-e267-d25c2c74397f', 'mvqq', 'Mvqq', true, '{"instagram":"https://www.instagram.com/andreas_bergeer/","twitter":"https://twitter.com/Mvqdh","youtube":"https://youtube.com/c/mvqdh"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '3ce4a46e-0ea4-36c8-e267-d25c2c74397f', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'havoc' AND c.class_slug = 'demon_hunter'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('cc96ac0f-8a98-6b46-7bd6-202560e7dd1c', 'trenacetatetv', 'Trenacetate', true, '{"discord":"https://discord.gg/agqdg2A","instagram":"http://instagram.com/jamaldinsworld/","twitter":"https://twitter.com/trenacetatetv","youtube":"https://www.youtube.com/channel/UCZwnwgvQxNX99RqfnBjlifw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'cc96ac0f-8a98-6b46-7bd6-202560e7dd1c', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'havoc' AND c.class_slug = 'demon_hunter'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('64f86317-2a59-bf81-6d85-c5956f3da4c6', 'razoonxz', 'Razoon', true, '{"youtube":"https://www.youtube.com/channel/UCSNjHLil0g8B4m2nlV_Pzfw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '64f86317-2a59-bf81-6d85-c5956f3da4c6', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'frost' AND c.class_slug = 'death_knight'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('403d3816-a9a3-e66c-5856-f2a22fe74974', 'gargoylemvp', 'Gargooyle', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '403d3816-a9a3-e66c-5856-f2a22fe74974', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'unholy' AND c.class_slug = 'death_knight'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('08d2127c-ce41-00f1-1f97-8909397c45a0', 'cerva', 'Cerva', true, '{"instagram":"https://www.instagram.com/cervatv/","twitter":"https://twitter.com/Cervatv","youtube":"https://www.youtube.com/channel/UCI2k32-vTf20BJHY6qnTB7w?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('4ec33e9b-9805-9dbd-d79e-b2ab7bf25fa0', 'losiro', 'Losiro', true, '{"twitter":"https://twitter.com/losirowow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('2bec3c0b-9806-628b-255b-fdeef3c0dec9', 'skippeetv', 'Skippee', true, '{"twitter":"https://twitter.com/skippeedk"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('c6da3267-f21c-53d7-afda-be8c284a2902', 'exyth1', 'Exyth', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'c6da3267-f21c-53d7-afda-be8c284a2902', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'unholy' AND c.class_slug = 'death_knight'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('814d0a53-f1c2-d65e-68f0-07bfb8c0f4e8', 'notmes', 'Mes', true, '{"twitter":"https://twitter.com/mes_wow","youtube":"https://www.youtube.com/channel/UC3yooAg-uF1xwa1OHQ1NK6g"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '814d0a53-f1c2-d65e-68f0-07bfb8c0f4e8', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'unholy' AND c.class_slug = 'death_knight'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '814d0a53-f1c2-d65e-68f0-07bfb8c0f4e8', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'arms' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('1ddac6ed-9936-1b05-d742-0075dc745539', 'bobydk1', 'Bobydk1', true, '{"discord":"https://discord.gg/FMmxmzZkQe","twitter":"https://twitter.com/bobydk1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('47ab37ba-2753-dcba-1c1b-c40200383c4f', 'creationlawlz', 'Creationlawlz', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '47ab37ba-2753-dcba-1c1b-c40200383c4f', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('04440f35-246b-d020-4956-8b4699d8b79a', 'xnatres', 'Xnatres', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '04440f35-246b-d020-4956-8b4699d8b79a', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'balance' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ef3981d7-286b-2db3-5242-c1a96e7555bc', 'simbo_feral', 'Simboferal', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'ef3981d7-286b-2db3-5242-c1a96e7555bc', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('b3fa55f9-8fcf-caf6-a15a-7c4eb7cdd1b5', 'bean', 'Bean', true, '{"twitter":"https://twitter.com/beantwitch","youtube":"https://youtube.com/beantwitch"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'b3fa55f9-8fcf-caf6-a15a-7c4eb7cdd1b5', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('6edab50a-9d0e-397c-2305-5e92f4d54667', 'shyxy_tv', 'Shyxy', true, '{"youtube":"https://www.youtube.com/c/ShyxyTV"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '6edab50a-9d0e-397c-2305-5e92f4d54667', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('110f7fa2-5690-3412-da75-bf5524583796', 'moonfirebeam', 'Moonfirebeam', true, '{"discord":"https://discord.gg/Z8TgJy6USv","twitter":"https://twitter.com/moonfirebeam","youtube":"https://www.youtube.com/channel/UCsSYn11a0R4wyZUHyd57yIA"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '110f7fa2-5690-3412-da75-bf5524583796', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'balance' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('98f4eb4e-403e-af45-1df4-f5e074e211e8', 'spottman', 'Spottman', true, '{"twitter":"https://twitter.com/SpottmanTV","youtube":"https://www.youtube.com/channel/UCQCZMWxDuofOStBWprVlK1Q"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '98f4eb4e-403e-af45-1df4-f5e074e211e8', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('fab1832a-d530-46d4-43a3-91fdbad89260', 'asgarathpvp', 'Asgarath', true, '{"twitter":"https://twitter.com/Asgarathpvp"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'fab1832a-d530-46d4-43a3-91fdbad89260', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('7dc9fd82-77f6-010d-d66c-c7ed5e5d286a', 'claak', 'Claak', true, '{"twitter":"https://twitter.com/claak_tv","youtube":"https://www.youtube.com/channel/UCPESiEiyz59_KfY6vemeo4g"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('9de91475-4c03-3593-d657-5b63c6fac40e', 'tonyferalmovies', 'Tonyferal', true, '{"discord":"https://discord.gg/443Xe3U","twitter":"https://twitter.com/tonyferalmovies","youtube":"https://www.youtube.com/c/Tonyferalmovies?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '9de91475-4c03-3593-d657-5b63c6fac40e', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('494b695c-f853-51dc-1cf6-90bc86685133', 'healingstat', 'Healingstat', true, '{"twitter":"https://twitter.com/Healingstat"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('c88baa9f-5fcd-0eff-c059-056c86466ab0', 'luuxia', 'Luuxia', true, '{"discord":"https://discord.gg/xtKy3pk","instagram":"https://www.instagram.com/luuxiaoff/?hl=fr","twitter":"https://twitter.com/JLuuxiA"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'c88baa9f-5fcd-0eff-c059-056c86466ab0', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'c88baa9f-5fcd-0eff-c059-056c86466ab0', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'c88baa9f-5fcd-0eff-c059-056c86466ab0', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('1de05c99-32ed-3e65-4248-e3c7f8449059', 'sodapoppin', 'Sodapoppin', true, '{"discord":"https://discord.gg/sodapoppin","instagram":"https://www.instagram.com/sodapoppintv/","twitter":"https://twitter.com/Sodapoppintv/","youtube":"https://www.youtube.com/c/sodapoppin"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '1de05c99-32ed-3e65-4248-e3c7f8449059', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('6519593b-4d0d-f0e8-7fb7-45158fee4498', 'snupy', 'Snupy', true, '{"discord":"https://discord.gg/h2yAPBB","reddit":"https://www.reddit.com/r/Snupy/","twitter":"https://twitter.com/snupytv","youtube":"https://www.youtube.com/channel/UC8epYaippj44sS41P3nLfkw/"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '6519593b-4d0d-f0e8-7fb7-45158fee4498', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'feral' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '6519593b-4d0d-f0e8-7fb7-45158fee4498', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'elemental' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('f6fa78c2-c019-02f2-3810-81303611a2c6', 'pszjager', 'Pszjager', true, '{"discord":"https://discord.gg/hxDy7FU","twitter":"https://twitter.com/pszjager","youtube":"https://www.youtube.com/channel/UCL-IFyMkjFeMU9Aiee1f8DQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('268f0570-00be-b4d0-a307-fd6a09c8cd76', 'haraw', 'Haraw', true, '{"twitter":"https://twitter.com/Harawlul","youtube":"https://www.youtube.com/HARAWW"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '268f0570-00be-b4d0-a307-fd6a09c8cd76', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'arms' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('67823369-ea3e-249e-3091-c5aaaea9fdbb', 'ilovelucy_', 'Ilovelucy', true, '{"twitter":"https://twitter.com/ilovelucy_wow","youtube":"https://www.youtube.com/playlist?list=PLg2w4zWZQhvc6opGac6_3OxqxyGQiSOuy"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('58de4a3f-c07a-3b5c-9934-bbc51bb34754', 'supatease', 'Supatease', true, '{"instagram":"https://www.instagram.com/supacasting/","twitter":"https://twitter.com/Supatease","youtube":"https://www.youtube.com/channel/UCt6INY6VJraOZBVWTJKmvkQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '58de4a3f-c07a-3b5c-9934-bbc51bb34754', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('c444541a-0b02-4f24-211a-86a0b8312d23', 'minpojke', 'Minpojke', true, '{"discord":"https://discord.gg/STTWtwwDet","instagram":"https://www.instagram.com/jonandersson92/","twitter":"https://twitter.com/Minpojke_","youtube":"https://www.youtube.com/minpojke"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'c444541a-0b02-4f24-211a-86a0b8312d23', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('c64a5413-dba0-448c-69f3-335067761bf9', 'bankstair', 'Bankstair', true, '{"youtube":"https://www.youtube.com/channel/UCwp2cj0mC7ws4PuF4AssweA"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('a957a7ea-68d0-830d-6211-1809c6852170', 'namesix1', 'Nayressa', true, '{"twitter":"https://twitter.com/tibort5"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'a957a7ea-68d0-830d-6211-1809c6852170', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'marksmanship' AND c.class_slug = 'hunter'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('c09a5303-23d1-4c5f-2029-a0924caf1aa4', 'homerjay_tv', 'Homerjay', true, '{"youtube":"https://www.youtube.com/channel/UCI8RZA6ZYtC_c0syNVUwqYg"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'c09a5303-23d1-4c5f-2029-a0924caf1aa4', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'marksmanship' AND c.class_slug = 'hunter'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('5ac239ec-45a2-5099-a6e7-41d8acea9653', 'paradise1_', 'Paradise1', true, '{"discord":"https://discord.gg/NR38xs7","twitter":"https://twitter.com/Paradise1WoW","youtube":"https://www.youtube.com/channel/UC3Ui8dwMLAY_GMrFqj5ztEw?"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('fabf4d77-91ee-65a0-c785-02f0c41f1dc0', 'Disyo', 'Disyo', true, '{"twitter":"https://twitter.com/Disyowow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('c92eab33-c275-a469-5236-95474545e850', 'kasuxoxo', 'Kasuxoxo', true, '{"discord":"https://discord.gg/5W3YWV7Hbn","instagram":"https://www.instagram.com/kasuhoho/","twitter":"https://twitter.com/kasuxoxoz","youtube":"https://www.youtube.com/channel/UCu5LyP1luajA3c8EZ6kgI7g?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('1aa523fc-d782-8a1b-7def-a781d700dded', 'billybobbs', 'Billybobbs', true, '{"youtube":"https://www.youtube.com/channel/UC7evJCfohlRtzNSY5QU3Mcg"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('7566fd9c-3a06-7581-7799-e3c99a452eeb', 'dilly_wow', 'Dilly', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e0e50e0b-2539-64a8-0269-1bbfc3337662', 'ssds', 'Ssds', true, '{"twitter":"https://twitter.com/SsdsLw","youtube":"https://www.youtube.com/channel/UC_6s9qDs2NZme0Rnc9rN2_g"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d1e28d14-f9dd-1934-32f7-fd708994afaf', 'jellybeans', 'Jellybeans', true, '{"instagram":"https://www.instagram.com/vincenttrangg/","twitter":"https://twitter.com/JellybeansTV","youtube":"https://www.youtube.com/channel/UCsfnT_8Dr5j3xpvz0azHmaA?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'd1e28d14-f9dd-1934-32f7-fd708994afaf', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'survival' AND c.class_slug = 'hunter'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('f41fcede-2712-32e6-3341-a705aa7f96b2', 'bicmexwow', 'Bicmexwow', true, '{"discord":"https://discord.gg/6Rdaz8K","twitter":"https://twitter.com/bicmexwow","youtube":"https://www.youtube.com/user/NullesWa"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('6998b93d-4f68-cdb6-e98d-5d1904ea9e10', 'freezion', 'Freezion', true, '{"discord":"https://discord.gg/gtuMn3f","youtube":"https://www.youtube.com/watch?v=gdGHz0GVXH8"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '6998b93d-4f68-cdb6-e98d-5d1904ea9e10', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'marksmanship' AND c.class_slug = 'hunter'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('bf245a2a-e541-e203-bce3-96c0373d8ccc', 'moopz', 'Moopz', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'bf245a2a-e541-e203-bce3-96c0373d8ccc', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'frost' AND c.class_slug = 'mage'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('68e6b461-0dd6-7b6f-f6e1-d256688ebd27', 'rivahlol', 'Rivah', true, '{"discord":"https://discord.gg/KfTrWvBur2","twitter":"https://twitter.com/RivahTBC","youtube":"https://www.youtube.com/channel/UCF7oCMiG304KSR1BbkutIzg?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '68e6b461-0dd6-7b6f-f6e1-d256688ebd27', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'frost' AND c.class_slug = 'mage'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('3e6dc3e2-bb4f-018c-b437-89c47a7c5555', 'cmdxyz', 'Cmdxyz', true, '{"twitter":"https://twitter.com/cmdzyx"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('1022c2c7-b58b-3a0b-3a5a-bc92d5cd32f4', 'pherix1', 'Pherix1', true, '{"discord":"https://discord.gg/hXVUb2","twitter":"https://twitter.com/DG_pherix"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('02916994-e72e-8657-8fa7-67b3e2f2a671', 'vultz', 'Vultz', true, '{"discord":"https://discord.gg/XTXgsRx","twitter":"https://twitter.com/Vultzover"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d281e9a5-40e9-0017-83d8-04a715a0389e', 'verinasty', 'Verinasty', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('eff0b205-f6fd-df56-1fc8-aec381d96759', 'samiyam', 'Samiyam', true, '{"twitter":"https://twitter.com/samiyamwow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('443631ac-be45-08a1-567a-c1cfa643923e', 'wealthymanx', 'Wealthyman', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('9da307e8-5a7f-3f6a-a640-460e02ed299d', 'wowalec', 'Alec', true, '{"youtube":"https://www.youtube.com/channel/UCBXOjQDF5Q96Qf0KNttBqow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('4723158a-149b-bfc8-74e6-8782e29daba5', 'hansol', 'Hansol', true, '{"discord":"https://discord.gg/vk82Gj6","twitter":"http://www.twitter.com/hansolonfire","youtube":"http://www.youtube.com/azncoolz"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4723158a-149b-bfc8-74e6-8782e29daba5', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'fire' AND c.class_slug = 'mage'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e4c66e8b-f17c-deb3-fbd5-4fb7f4a3bc84', 'draizn', 'Draizn', true, '{"youtube":"https://www.youtube.com/channel/UC5b9NGeDXyBTg-RVjNhm47A"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('3c9acc1a-afca-c634-c249-dc5c60ed4f8a', 'albraik', 'Albraik', true, '{"twitter":"https://twitter.com/Albraiik"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '3c9acc1a-afca-c634-c249-dc5c60ed4f8a', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'frost' AND c.class_slug = 'mage'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('94567184-7e0c-bfb0-c942-5e96d99e8680', 'venruki', 'Venruki', true, '{"instagram":"https://www.instagram.com/venczel/","twitter":"https://twitter.com/ElliottVenczel","youtube":"https://www.youtube.com/venruki"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '94567184-7e0c-bfb0-c942-5e96d99e8680', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'fire' AND c.class_slug = 'mage'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('84a2c7dd-b388-7a28-6705-c0a58ad2c232', 'ziqoftw', 'Ziqoftw', true, '{"instagram":"https://www.instagram.com/ziqoftw/","twitter":"https://www.twitter.com/ziqoftw","youtube":"http://www.youtube.com/Ziqoftw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d6ae7487-7358-211b-00fb-a2ad9fd89f21', 'zqitv', 'Zqitv', true, '{"instagram":"https://www.instagram.com/tarz6n/","youtube":"https://www.youtube.com/user/Muffintoppinlolz/videos?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'd6ae7487-7358-211b-00fb-a2ad9fd89f21', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'frost' AND c.class_slug = 'mage'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('b9adb38c-0ca3-7443-1e43-5027bc3ba55b', 'raikubest', 'Raikubest', true, '{"twitter":"https://twitter.com/Raiku_Wow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'b9adb38c-0ca3-7443-1e43-5027bc3ba55b', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'frost' AND c.class_slug = 'mage'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('4952eea2-4041-bc6d-f3d0-8123b767cec0', 'xaryu', 'Xaryu', true, '{"discord":"https://discord.gg/xaryu","instagram":"https://www.instagram.com/joshlujan/","twitter":"https://twitter.com/Xaryu","website":"https://xaryu.tv/","youtube":"https://www.youtube.com/xaryu?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('8223fdd9-2496-af3d-8140-1d0dab5872b0', 'mysticallx', 'Mysticallx', true, '{"twitter":"https://twitter.com/MysticalTheMonk"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '8223fdd9-2496-af3d-8140-1d0dab5872b0', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'mistweaver' AND c.class_slug = 'monk'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('db73fd02-b988-5ead-b9c0-76f2612cd029', 'panzerhenk', 'Panzerhenk', true, '{"instagram":"https://www.instagram.com/panzerhenk/","twitter":"https://twitter.com/panzerhenk","youtube":"https://www.youtube.com/watch?v=dQw4w9WgXcQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'db73fd02-b988-5ead-b9c0-76f2612cd029', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'mistweaver' AND c.class_slug = 'monk'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('51ca5007-1aa8-cdbc-e0c1-95b795a16ad9', 'banwellx', 'Banwellx', true, '{"discord":"https://discord.gg/WtFZvrA","twitter":"https://www.twitter.com/banwell","youtube":"https://www.youtube.com/c/banwellx"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('97cbb7a5-c7c4-5c3d-1448-2774dd05a92e', 'woopy', 'Woopy', true, '{"twitter":"https://twitter.com/Woopywow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d88f4265-a03d-b8f7-51f7-d104cfa6a622', 'chunliww', 'Chunliww', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('dea72bf7-7534-4db5-06bc-0f39dcf6ed38', 'trilltko', 'Trilltko', true, '{"twitter":"https://twitter.com/TrilltkoWW"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('b3d48cdb-1b74-aeb1-880d-ef1aa0cf579e', 'orbrexth', 'Orbrexth', true, '{"twitter":"https://twitter.com/orbreXth","youtube":"https://www.youtube.com/c/orbreXth/videos"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'b3d48cdb-1b74-aeb1-880d-ef1aa0cf579e', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'mistweaver' AND c.class_slug = 'monk'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'b3d48cdb-1b74-aeb1-880d-ef1aa0cf579e', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('8868fb2d-6db7-469f-7ad3-94479442c468', 'trillebartom', 'Trillebartom', true, '{"discord":"https://discord.gg/trille","twitter":"https://twitter.com/trillebartom","youtube":"https://www.youtube.com/user/trillebartom?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d4f252c0-2289-bf8d-c4ba-a3bc6802a48f', 'h0lyundead', 'H0lyundead', true, '{"discord":"https://discord.gg/BxgGkZ7nEE","youtube":"https://www.youtube.com/channel/UCIxXLM7vT8JKc-UF9cJ9RTw?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'd4f252c0-2289-bf8d-c4ba-a3bc6802a48f', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('3c6f5923-1fe2-3d47-f6d7-caeaa2b6a5d7', 'earpugs', 'Earpugs', true, '{"youtube":"https://www.youtube.com/channel/UCTxM4AZCZKfyxNH8l2hm4uQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '3c6f5923-1fe2-3d47-f6d7-caeaa2b6a5d7', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('b459fd1f-3d46-0811-7d37-af0040477abc', 'tempestxgg', 'Tempest', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'b459fd1f-3d46-0811-7d37-af0040477abc', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ed662a8a-78e6-a98d-4c49-7ea4a8b7998c', 'ascalontv', 'Ascalontv', true, '{"discord":"https://discord.com/invite/uepKAsY","twitter":"https://twitter.com/AscalonTV","youtube":"https://www.youtube.com/channel/UC6FKFoJWGtZoKIDZRSXg78A"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'ed662a8a-78e6-a98d-4c49-7ea4a8b7998c', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'shadow' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('444f7f30-e460-73be-3065-a5eddb3ad5c1', 'hozitojones', 'Hozitojones', true, '{"discord":"https://discord.gg/V5eNAW5","instagram":"https://www.instagram.com/sebas_toriello/?hl=es-la","twitter":"https://twitter.com/Hozitojones","youtube":"https://www.youtube.com/channel/UC1cd0NFWO8XxgLhg23F1y-w?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '444f7f30-e460-73be-3065-a5eddb3ad5c1', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('4c0d77f2-5e56-0fea-9e31-a7799680cae5', 'aphobiagaming', 'Aphobia', true, '{"instagram":"https://www.instagram.com/eddphobia91/","youtube":"https://www.youtube.com/user/eddphobia"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4c0d77f2-5e56-0fea-9e31-a7799680cae5', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'shadow' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('02617310-3085-9586-2925-deae1989a994', 'olliektv', 'Olliektv', true, '{"discord":"https://discord.gg/GbQrSUyR"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '02617310-3085-9586-2925-deae1989a994', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('cfca0dff-d69f-0fef-56b3-41377bbe6354', 'stahpsp', 'Stahpsp', true, '{"youtube":"https://www.youtube.com/stahpsp"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'cfca0dff-d69f-0fef-56b3-41377bbe6354', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'shadow' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('78d57cf4-2764-fd49-63ae-0784473fc298', 'krawnzlol', 'Krawnzlol', true, '{"instagram":"https://www.instagram.com/jraranda7"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '78d57cf4-2764-fd49-63ae-0784473fc298', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d6c2eeee-f97a-9935-9c99-02054b0a5367', 'wizkx', 'Wizkx', true, '{"twitter":"https://twitter.com/wizkxd"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('f7c51bbf-5315-97ba-339f-ad99f7aef59e', 'vilaye', 'Vilaye', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('b15d1b10-b2cc-2be2-2982-ff98228a6500', 'hydramist', 'Hydramist', true, '{"facebook":"http://www.facebook.com/pages/Hydramist/147296701974820?ref=hl","twitter":"https://twitter.com/hydramist","youtube":"https://youtube.com/hydramist"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'b15d1b10-b2cc-2be2-2982-ff98228a6500', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('4635f403-3048-83be-bd92-8d2bc3100ebe', 'zenlyn', 'Zenlyn', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('fa1dfd54-4077-1168-c8a8-eddf1ee01d38', 'anboniwow', 'Anboniwow', true, '{"discord":"https://discord.gg/EhKKHEm","twitter":"https://www.twitter.com/anboniwow","youtube":"https://www.youtube.com/anboniwow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'fa1dfd54-4077-1168-c8a8-eddf1ee01d38', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'shadow' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('05503366-663c-77d4-03ee-dbcb81cfb4c1', 'mehhx', 'Mehhx', true, '{"twitter":"https://twitter.com/JuhaniHalme"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '05503366-663c-77d4-03ee-dbcb81cfb4c1', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'discipline' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('453be20a-e0d0-88c6-7f47-b920bbe794a0', 'gekz', 'Gekz', true, '{"discord":"https://discord.gg/6M6S8RBhKR","twitter":"https://twitter.com/Gekzs"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('bee1e838-f055-214f-1d32-c68ea8ba012f', 'chastv', 'Chastv', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('93063162-68d4-a56f-f1a9-aea3018f1693', 'Eltoni52', 'Eltoni52', true, '{"discord":"https://discord.gg/fcxQfrYe5r","instagram":"https://www.instagram.com/eltoni522/","twitter":"https://twitter.com/eltoni25","youtube":"https://www.youtube.com/channel/UCPwd-YmkC1aZGX7Gk91ucZw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('2e8dbf7d-6cb2-7c83-add5-7638bed1dec4', 'chokopapatv', 'Chokopapatv', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '2e8dbf7d-6cb2-7c83-add5-7638bed1dec4', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('95353955-dc44-28e6-479e-102359c5f44b', 'preghierax', 'Preghierax', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '95353955-dc44-28e6-479e-102359c5f44b', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'retribution' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('bd1650ca-02a8-1e70-0157-7917023eedea', 'flowstateswow', 'Flowstateswow', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'bd1650ca-02a8-1e70-0157-7917023eedea', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('f0d800a9-084d-d9c0-e20c-1dfb9a495aac', 'pizzahunter2009', 'Holysaxyboy', true, '{"discord":"https://discord.gg/W8M5NMF","youtube":"https://www.youtube.com/channel/UC1yMtI5eIE4cTex7KTfxb0w"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'f0d800a9-084d-d9c0-e20c-1dfb9a495aac', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'retribution' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('6c686338-4c27-1a73-c3b7-b1ad71fe942e', 'waynxt', 'Wayne', true, '{"youtube":"https://www.youtube.com/user/BarneyGamingCZ/videos"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '6c686338-4c27-1a73-c3b7-b1ad71fe942e', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('eeb0334c-7aea-f12a-4916-00a215c19e0b', 'flyntv_', 'Saori', true, '{"discord":"https://discord.gg/SSJzzzupBQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'eeb0334c-7aea-f12a-4916-00a215c19e0b', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('b1aa7e73-cdc3-dd54-f54a-bd6d64f8391b', 'soumi', 'Soumi', true, '{"discord":"https://discord.gg/3882XAKwyg"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'b1aa7e73-cdc3-dd54-f54a-bd6d64f8391b', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('c38a84b5-f052-2f3e-cd94-b760512156af', 'abxtv', 'abxtv', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d0b2b460-8b5c-7fe0-527e-fe366246d058', 'crusader3455', 'Crusader', true, '{"instagram":"https://www.instagram.com/itsmikeowens/","twitter":"https://twitter.com/itsmikeowens","youtube":"https://www.youtube.com/channel/UCzoNEIaIy1Q-VrrU1P-5vpw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'd0b2b460-8b5c-7fe0-527e-fe366246d058', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'retribution' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('89867ea7-5082-c9d2-2d5b-ac90dd0caefd', 'borngood', 'Borngood', true, '{"twitter":"https://twitter.com/BorngoodW"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('10432b75-9770-c485-0e00-92f1f9b4e617', 'rubcub', 'Rubcub', true, '{"twitter":"https://twitter.com/PG_Rubcub"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('68e74dec-5646-122a-88de-19cbed108122', 'savix', 'Savix', true, '{"discord":"https://discord.gg/savix","instagram":"https://www.instagram.com/savixirl/","reddit":"https://www.reddit.com/r/Savix/","twitter":"https://twitter.com/SavixIrL","youtube":"https://www.youtube.com/channel/UCIRe2YighqgPmSDjSb3Fpiw?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '68e74dec-5646-122a-88de-19cbed108122', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'retribution' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('9da0a1eb-3f33-7db0-a1a6-3df455b4daad', 'pvplab', 'Pvplab', true, '{"discord":"http://discord.gg/pvplab","instagram":"https://www.instagram.com/pvplablive/","twitter":"http://twitter.com/pvplablive"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '9da0a1eb-3f33-7db0-a1a6-3df455b4daad', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('4bd32bb6-5ceb-abea-98a3-c5b910e33a48', 'accident_', 'Accident', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4bd32bb6-5ceb-abea-98a3-c5b910e33a48', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'retribution' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('8dae1856-5127-f838-4089-df45f623a62c', 'bushy25', 'Bushy25', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('67b23557-fda6-583a-f95b-444d3ccf695e', 'snowmixy', 'Snowmixy', true, '{"instagram":"https://www.instagram.com/snowmixy/","twitter":"https://twitter.com/Snowmixy","youtube":"https://www.youtube.com/user/Snowmixy"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('fb305b8c-ec4a-0c09-0017-5f0b52b0ba03', 'vanguardstv', 'Vanguards', true, '{"discord":"https://discord.gg/Vanguards","twitter":"https://twitter.com/VanguardsTV","youtube":"https://www.youtube.com/c/Vanguards"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('03b99634-e494-7174-0a8b-aa23263eb8c2', 'pika_pala', 'Pikachu', true, '{"twitter":"https://twitter.com/pala_pika"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('68bf225d-7db8-3f75-046d-85b8676b15d0', 'zabius', 'Zabius', true, '{"twitter":"https://twitter.com/FelixFahauer1","youtube":"https://www.youtube.com/channel/UCZSfktks-roK7sdh6DmdgMQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('56a20cd6-f154-8eb5-6ba1-f171a1e715cb', 'reqtbc', 'Reqtbc', true, '{"discord":"https://discord.gg/DVQYXk4"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '56a20cd6-f154-8eb5-6ba1-f171a1e715cb', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'subtlety' AND c.class_slug = 'rogue'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('45d780f8-2c92-47fc-1142-f8b60c12b6c6', 'perplexity', 'Perplexity', true, '{"discord":"https://discord.gg/aXryrUK","youtube":"https://www.youtube.com/watch?v=fLkZuz7kjyU\u0026ab_channel=perplexity"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '45d780f8-2c92-47fc-1142-f8b60c12b6c6', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'subtlety' AND c.class_slug = 'rogue'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e653ad4a-a4bb-f0d8-9e4a-f865d4e28caf', 'avizura', 'Avizura', true, '{"twitter":"https://twitter.com/avizuray","youtube":"https://www.youtube.com/c/Avizura/featured?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e653ad4a-a4bb-f0d8-9e4a-f865d4e28caf', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'subtlety' AND c.class_slug = 'rogue'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ec4f758a-0079-e9b1-ff74-300197f910e7', 'petraxs', 'Petraxs', true, '{"youtube":"https://www.youtube.com/c/Petraxs"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'ec4f758a-0079-e9b1-ff74-300197f910e7', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'subtlety' AND c.class_slug = 'rogue'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('051fd026-ba83-ae50-850b-c815ef354505', 'traxelol', 'Traxelol', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('975de47b-99b8-632a-ac2c-4b1d6059e0b7', 'jiberjaber', 'Jiberjaber', true, '{"discord":"https://discord.gg/PfqC5gP7Mp"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('a7cbfbaf-974b-43fb-7cb0-63f57ad913e4', 'chickchau', 'Chickchau', true, '{"discord":"https://discord.gg/jX7rfQHUZX","instagram":"https://www.instagram.com/chick_chau/"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ced0114e-fb52-befb-7b93-087b37e128d7', 'NAIIKAII', 'Naiikaii', true, '{"discord":"https://discord.gg/eEMUwuS","instagram":"https://www.instagram.com/naiikaii_/","twitter":"https://twitter.com/NaiikaiiGG","youtube":"https://www.youtube.com/c/NAIIKAII?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('bca65e59-c795-1ac8-e21c-93d046f78d5f', 'nahj', 'Nahj', true, '{"twitter":"https://twitter.com/RogueNahj","youtube":"https://www.youtube.com/nahjwow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('a4bb651e-a626-b53a-2fa7-30d7c61d4c4c', 'mirlol', 'Mirlol', true, '{"twitter":"https://twitter.com/mirlol","youtube":"https://www.youtube.com/channel/UC3r7qQGEPcRuHGmMNCZdS4Q"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'a4bb651e-a626-b53a-2fa7-30d7c61d4c4c', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'subtlety' AND c.class_slug = 'rogue'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d09f9c80-0bd2-30cf-44ec-204a6b470ee9', 'palumor', 'Palumor', true, '{"discord":"https://discord.gg/ydJ5eUpZWD","twitter":"https://twitter.com/iPalumor","youtube":"https://www.youtube.com/c/Palumor"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ca87a399-b549-18b1-f803-54db4c1b09ce', 'pikabooirl', 'Pikaboo', true, '{"instagram":"https://www.instagram.com/jasonsmithtm/","twitter":"https://twitter.com/pikaboowow","youtube":"https://www.youtube.com/channel/UCrYo6np138Ge6iCRMgS93aQ?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e1705f6b-00f0-3c11-5aef-65a5c9492ff9', 'stungodx', 'Stungodx', true, '{"twitter":"https://twitter.com/Stungodx"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('adc45882-8ba5-691d-1e92-fb9844e39802', 'aidenzx', 'Aidenzx', true, '{"twitter":"https://twitter.com/UN_Aiden"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('cbd62b8e-6ec8-bd0c-509b-1bcf3fbd641d', 'SensusLive', 'Sensuslive', true, '{"discord":"https://discord.gg/jzJ3P6eEGr","twitter":"https://twitter.com/SensusYT","youtube":"https://www.youtube.com/SensusYT"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('785dbb72-393e-e1c0-b1e8-4be6222bee12', 'jaxington', 'Jaxington', true, '{"discord":"https://discord.gg/jExCyhv","twitter":"https://twitter.com/jaxirl","youtube":"https://youtube.com/skidsh"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('0f45a71c-e60f-ce95-07c6-dc7620ac7e24', 'akrololz', 'Akrololz', true, '{"discord":"https://discord.gg/akrololz","instagram":"https://www.instagram.com/akrobro/","twitter":"https://twitter.com/Akrololz","youtube":"https://www.youtube.com/akrololz?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('27f6ff91-37a2-9610-59a5-6e685a480e62', 'whaazz', 'Whaazz', true, '{"instagram":"https://www.instagram.com/oscar_wulff/","twitter":"https://twitter.com/Whaazz","youtube":"https://www.youtube.com/channel/UCJFHadK6prEPYQuCmR0tsug"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('faff007a-1f9e-b7a0-30bc-6d14b6ee7a9f', 'psherotv', 'Pshero', true, '{"discord":"https://discord.gg/h7nrY8j","instagram":"https://www.instagram.com/psherotv/","twitter":"https://twitter.com/Psherotv","youtube":"https://www.youtube.com/psheor"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'faff007a-1f9e-b7a0-30bc-6d14b6ee7a9f', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'subtlety' AND c.class_slug = 'rogue'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('49fda3a4-fd55-c1ee-e2ef-92f7cb5e7f60', 'knoffwow', 'Knoffwow', true, '{"discord":"https://discord.gg/atKfD52hfm","instagram":"https://www.instagram.com/knoffwow/","twitter":"https://www.twitter.com/knoffwow","youtube":"https://www.youtube.com/knoffwow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e84ae83f-7d28-f839-3804-4365c4a777e8', 'frozentherogue', 'Frozen', true, '{"website":"https://www.donationalerts.com/r/frozentherogue"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e84ae83f-7d28-f839-3804-4365c4a777e8', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'assassination' AND c.class_slug = 'rogue'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('531240ef-1c17-40cc-3228-34014cb97921', 'groblinwow', 'Groblin', true, '{"discord":"https://discord.gg/SFJ8eNnDYS"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '531240ef-1c17-40cc-3228-34014cb97921', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('8d93315a-8726-7225-7993-c410b164886c', 'zakiwak1', 'Zakiwaki', true, '{"discord":"https://discord.gg/8akXAsnQ4p","youtube":"https://www.youtube.com/channel/UCaXjkbB_eCUpQo76raxYa1A"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '8d93315a-8726-7225-7993-c410b164886c', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('47dd23b7-9fbd-a938-2ce2-064362edf6b5', 'traktorch0', 'Traktorcho', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '47dd23b7-9fbd-a938-2ce2-064362edf6b5', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'enhancement' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('de3abe77-1e43-f5f4-9f83-abeb2890b9aa', 'skilerxtv', 'Skiler', true, '{"discord":"https://discord.gg/yewtFpej6j"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'de3abe77-1e43-f5f4-9f83-abeb2890b9aa', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'elemental' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('7fdd4909-8d0d-0891-093d-d939d91ef78f', 'blackbettytv', 'Blackbetty', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '7fdd4909-8d0d-0891-093d-d939d91ef78f', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'elemental' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e77e2e9f-3662-e46d-739f-df3c662227fe', 'novoz', 'Novoz', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e77e2e9f-3662-e46d-739f-df3c662227fe', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'elemental' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('80bb628b-5d39-af1d-7505-328cd1484493', 'hrkzz', 'Hrkzz', true, '{"instagram":"https://www.instagram.com/rn.fitt/?hl=fr"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '80bb628b-5d39-af1d-7505-328cd1484493', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'elemental' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('203649f2-b9aa-5061-8320-ef1e682986df', 'warbla', 'Warbla', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '203649f2-b9aa-5061-8320-ef1e682986df', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('3b1f6f09-7aa5-6c34-ec32-3d93421d3d1b', 'kolowavex', 'Kolowavex', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e7e54bbb-0564-aa8f-41f6-886896a719fb', 'tiqqlethis', 'Tiqqlethis', true, '{"discord":"https://discord.gg/gsky4Ykcg8","instagram":"https://www.instagram.com/tiqqleofficial/","twitter":"https://twitter.com/TiqqleThis","youtube":"https://www.youtube.com/tiqqle?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e7e54bbb-0564-aa8f-41f6-886896a719fb', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'enhancement' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('5fef297d-b6ff-becc-9b97-c10247d285c3', 'swapxy', 'Swapxy', true, '{"twitter":"https://twitter.com/Swapxy"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '5fef297d-b6ff-becc-9b97-c10247d285c3', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'enhancement' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('4a010dd9-c4f7-9652-9ba0-6b16297a0177', 'drainerx', 'Drainerx', true, '{"discord":"https://discord.gg/DMUSWhG","twitter":"https://twitter.com/intent/user?screen_name=Drainerxtv","youtube":"https://www.youtube.com/drainerxtv?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4a010dd9-c4f7-9652-9ba0-6b16297a0177', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4a010dd9-c4f7-9652-9ba0-6b16297a0177', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4a010dd9-c4f7-9652-9ba0-6b16297a0177', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4a010dd9-c4f7-9652-9ba0-6b16297a0177', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4a010dd9-c4f7-9652-9ba0-6b16297a0177', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'mistweaver' AND c.class_slug = 'monk'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '4a010dd9-c4f7-9652-9ba0-6b16297a0177', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'preservation' AND c.class_slug = 'evoker'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('1bc65594-08e8-7d48-fc32-44c8597b82bf', 'zeepeye', 'Zeepeye', true, '{"discord":"https://discord.gg/tCq9TQGaQm","twitter":"https://twitter.com/ZeepeyeSwE","youtube":"https://www.youtube.com/channel/UCmRgHhywwS7fp97tki9eY6g"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '1bc65594-08e8-7d48-fc32-44c8597b82bf', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'elemental' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('e35e8721-c542-d971-c135-aaa3a0362426', 'lontartv', 'Lontar', true, '{"instagram":"https://www.instagram.com/gabri_cano/","twitter":"https://twitter.com/Lontar_wow","youtube":"https://www.youtube.com/channel/UCCDFA0pRgHYZPx3l_uCRvAw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e35e8721-c542-d971-c135-aaa3a0362426', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'mistweaver' AND c.class_slug = 'monk'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e35e8721-c542-d971-c135-aaa3a0362426', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e35e8721-c542-d971-c135-aaa3a0362426', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'restoration' AND c.class_slug = 'druid'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e35e8721-c542-d971-c135-aaa3a0362426', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'paladin'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'e35e8721-c542-d971-c135-aaa3a0362426', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'holy' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('80649308-5d1b-8368-f1d6-736684f50cf9', 'cdewx', 'Cdew', true, '{"twitter":"https://twitter.com/cdew_wow","youtube":"https://www.youtube.com/user/tenderloinx1?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('82e4ffe8-cf2e-2e01-81bf-d44acc30515e', 'Lovelesstv', 'Loveless', true, '{"discord":"https://discord.gg/eFrG6fmM","twitter":"https://twitter.com/lovelesswow","youtube":"https://www.youtube.com/channel/UCsLbmQM0iqdM5_Fg2jRh3Gg"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d8890241-111e-a057-25fa-98d1bb4eab1d', 'infernion', 'infernion', true, '{"discord":"https://discord.gg/R8pKuhhba3","instagram":"https://www.instagram.com/jespergjessen/","twitter":"https://twitter.com/Infernionwow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'd8890241-111e-a057-25fa-98d1bb4eab1d', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'affliction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('178566f5-eb7f-dce4-c0f5-56a343708087', 'chanx', 'chanx', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '178566f5-eb7f-dce4-c0f5-56a343708087', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ad19ed73-ffa8-cde8-e42f-b9eb0c7ccf1e', 'kerosineyo', 'Kerosineyo', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'ad19ed73-ffa8-cde8-e42f-b9eb0c7ccf1e', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('476f695b-4198-e423-1f6a-175d4a091b04', 'scryna', 'Scryna', true, '{"twitter":"https://twitter.com/scryna","youtube":"https://www.youtube.com/@scryna"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '476f695b-4198-e423-1f6a-175d4a091b04', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('0b9f6622-54a4-7d45-1f87-dfc3ef9fd286', 'reverencewarlock', 'Reverence', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '0b9f6622-54a4-7d45-1f87-dfc3ef9fd286', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('a1b63e4b-dff8-536e-03bf-a62b09080f35', 'bualock', 'Bualock', true, '{"discord":"https://discord.gg/sxhwe3b","instagram":"https://instagram.com/BuaLockOW","twitter":"https://twitter.com/BuaLockOW","youtube":"https://www.youtube.com/c/bualock/videos"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'a1b63e4b-dff8-536e-03bf-a62b09080f35', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('afb3b9d3-dc42-7728-c58c-28b35b461de6', 'snutzy', 'Snutzy', true, '{"discord":"https://discord.gg/eMTkkB3","instagram":"https://www.instagram.com/kelnguyenn/","twitter":"https://twitter.com/Snutztv"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'afb3b9d3-dc42-7728-c58c-28b35b461de6', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'afb3b9d3-dc42-7728-c58c-28b35b461de6', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'shadow' AND c.class_slug = 'priest'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ef05d425-53e1-3d43-e00c-12f2d2de62a9', 'wallirik', 'Wallirik', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('cf79e0df-638e-613b-483c-23219550a2b6', 'jazggz', 'Jazggz', true, '{"discord":"https://discord.gg/ZfNg2n44WK","twitter":"https://twitter.com/jazggz","youtube":"https://www.youtube.com/channel/UCXWELAuYNuRk53qgIWPTFaA"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'cf79e0df-638e-613b-483c-23219550a2b6', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('1b472cc5-c403-c930-5dbf-f8b188684c8c', 'anniefuchsia', 'Anniefuchsia', true, '{"instagram":"https://www.instagram.com/anniefuchsia/","twitter":"https://twitter.com/anniefuchsia","youtube":"https://www.youtube.com/anniefuchsia?sub_confirmation=1"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '1b472cc5-c403-c930-5dbf-f8b188684c8c', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'destruction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('26ee8e32-a9f2-a7e6-46c2-0e38f503641f', 'maldiva', 'Maldiva', true, '{"discord":"https://discord.gg/maldiva","instagram":"https://www.instagram.com/maldivawow/","twitter":"https://twitter.com/Maldivawow","youtube":"https://www.youtube.com/user/Maldivapvp"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '26ee8e32-a9f2-a7e6-46c2-0e38f503641f', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'affliction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('d086f191-0b27-96b3-8a0d-a2c5732b9fdf', 'germinouu066', 'Germi', true, '{"discord":"https://discord.gg/RHDBxsN","twitter":"https://twitter.com/Germinou06"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('be7fc846-9a7b-b681-b156-2bdee8c17140', 'featherfeeet', 'Featherfeeet', true, '{}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('70f225a2-a63c-4544-4eaf-81073f7fff59', 'gelubabatv', 'Gelubabatv', true, '{"twitter":"https://twitter.com/CalGelu","youtube":"https://www.youtube.com/channel/UCskmJY6W20-v27tuIr5eDkA"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '70f225a2-a63c-4544-4eaf-81073f7fff59', s.id, false
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'elemental' AND c.class_slug = 'shaman'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '70f225a2-a63c-4544-4eaf-81073f7fff59', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'affliction' AND c.class_slug = 'warlock'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('96378ad4-0e66-d031-45eb-4e20159d87c1', 'mercedesa', 'Mercedesa', true, '{"discord":"https://discord.gg/Mau5mWQ","twitter":"https://twitter.com/mercewow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('8d17f835-0b49-08de-4440-9640fabf3cd8', 'dakkroth', 'Dakkroth', true, '{"facebook":"https://www.facebook.com/Dakkroth","twitter":"https://twitter.com/Dakkrothwow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('763dbc5b-676f-1bbe-08c4-1b3cac4a80e7', 'kostickatv', 'Kosticka', true, '{"discord":"https://discord.gg/ejWtuhR","youtube":"https://www.youtube.com/channel/UCT0QT0ebAi1U0myAwuoowvQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '763dbc5b-676f-1bbe-08c4-1b3cac4a80e7', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'protection' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('8799cc7f-c001-6f02-12ff-5e480ea2bdee', 'bcxwarr', 'Bcw', true, '{"discord":"https://discord.gg/PwCBmNeGDF"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '8799cc7f-c001-6f02-12ff-5e480ea2bdee', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'arms' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('36a40b21-183c-dea2-fb09-28d9a1a1e7db', 'torstenstock', 'Torstenstock', true, '{"discord":"https://discord.gg/jDVkDEZA"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '36a40b21-183c-dea2-fb09-28d9a1a1e7db', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'arms' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('ea509f96-6221-7504-ef60-2ddfcb566c43', 'pvelordtv', 'Pvelordtv', true, '{"youtube":"https://www.youtube.com/channel/UCUQMAH2NXCM1GKznX5QYlMQ"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('135da104-fa77-dab3-75f2-b11c6c5ea0a4', 'zhreytv', 'Zhreytv', true, '{"twitter":"https://twitter.com/Zhreytv"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('a0fe4222-c8c9-e2cf-a7fd-8f10071ce397', 'blizo', 'Blizo', true, '{"twitter":"https://twitter.com/blizowow"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('8fd65e3b-b971-7854-9da8-a3be3b440ffc', 'smexxin', 'Smexxin', true, '{"twitter":"https://twitter.com/Smexxin"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('62a56d74-02a5-f302-9323-6a3c0ed57c82', 'mitu_tv', 'Mitu', true, '{"twitter":"https://twitter.com/Mitu_TV","youtube":"https://www.youtube.com/channel/UC85rl8ND2SNxO3LC_tCTztw?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('2f625e9f-f919-a619-79b0-d987f07d237d', 'hito', 'Hito', true, '{"discord":"https://discord.com/invite/CNj5bjnBB8","twitter":"https://twitter.com/Hitowow","youtube":"https://www.youtube.com/c/Hitokiro"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT '2f625e9f-f919-a619-79b0-d987f07d237d', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'arms' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('cdd745f2-5c9e-3d3f-5ac0-0bcc14c21031', 'joefernandes123', 'Joefernandes', true, '{"discord":"https://discord.gg/MMpmYK3F","instagram":"https://www.instagram.com/joefernandes123/","twitter":"https://twitter.com/joefernandes123","youtube":"https://www.youtube.com/channel/UCRU4UWnsYB-zhaKkaF37EEw"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('415b243c-4861-03d2-c887-2b2ef0e70dc0', 'bajheera', 'Bajheera', true, '{"discord":"https://discord.gg/DWc3mFmEyd","instagram":"https://www.instagram.com/bajheerawow/","twitter":"https://twitter.com/BajheeraWoW","youtube":"https://www.youtube.com/user/BajheeraWoW"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('bd949994-b373-51a4-6311-2826a13d26aa', 'dekel', 'Dekel', true, '{"discord":"https://discord.gg/jPb9tr9","instagram":"https://instagram.com/dekeldeclan","twitter":"https://twitter.com/dekeldeclan","youtube":"https://www.youtube.com/channel/UCMEwlq2-4FzgzeGllkRHJSA?view_as=subscriber"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'bd949994-b373-51a4-6311-2826a13d26aa', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'arms' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
INSERT INTO public.channels (id, twitch_login, display_name, is_published, socials) VALUES ('dd130d38-40ae-061a-bfdc-0fce945135b1', 'magnusz', 'Magnusz', true, '{"discord":"https://discord.gg/dFtqrHe","twitter":"https://twitter.com/sx6_magnusz","youtube":"https://www.youtube.com/magnusz"}') ON CONFLICT (twitch_login) DO UPDATE SET display_name = EXCLUDED.display_name, socials = EXCLUDED.socials;
INSERT INTO public.channel_specs (channel_id, spec_id, is_primary)
SELECT 'dd130d38-40ae-061a-bfdc-0fce945135b1', s.id, true
FROM specs s
JOIN classes c ON s.class_id = c.id
WHERE s.spec_slug = 'arms' AND c.class_slug = 'warrior'
ON CONFLICT (channel_id, spec_id) DO UPDATE SET is_primary = EXCLUDED.is_primary;
