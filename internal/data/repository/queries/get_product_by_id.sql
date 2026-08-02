SELECT 
    p.id,
    p.category_id,
    p.title_en,
    p.title_fa,
    p.description_en,
    p.description_fa,
    p.price_toman,
    p.price_usd,
    p.previous_price_toman,
    p.previous_price_usd,
    p.discount_percentage,
    p.images,
    p.created_at,
    c.slug
FROM products p 
LEFT JOIN categories c ON p.category_id = c.id  
WHERE p.id = $1