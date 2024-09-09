package jtable

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"db/disasterdb"
	"db/common"
)

// TODO
// 1 translate region
// 2 convert time

func lookupLocation(reg string) string {
	return map[string]string {
		"Средиземное море": "41.29254,12.573465",
		"Троя": "41.29254,12.573465",
		"Испания \n(Кадисский залив)": "41.29254,12.573465",
		"Греция": "41.29254,12.573465",
		"Япония": "41.29254,12.573465",
		"Египет": "41.29254,12.573465",
		"Африка": "41.29254,12.573465",
		"Финикия": "41.29254,12.573465",
		"Китай": "41.29254,12.573465",
		"Сирия": "41.29254,12.573465",
		"Керченский пролив": "41.29254,12.573465",
		"Палестина": "41.29254,12.573465",
		"Иудея": "41.29254,12.573465",
		"Побережье \nЧерного моря": "41.29254,12.573465",
		"Италия": "41.29254,12.573465",
		"О. Кипр (Крит)": "41.29254,12.573465",
		"Южная Турция": "41.29254,12.573465",
		"Малая Азия, Рим": "41.29254,12.573465",
		"Турция": "41.29254,12.573465",
		"Малая Азия, \nСредиземное море": "41.29254,12.573465",
		"Палестина, Сирия": "41.29254,12.573465",
		"Восточное Средиземноморье": "41.29254,12.573465",
		"Малая Азия": "41.29254,12.573465",
		"Мраморное море": "41.29254,12.573465",
		"Византия, Малая Азия, Междуречье": "41.29254,12.573465",
		"Югославия": "41.29254,12.573465",
		"Сирия, Малая Азия": "41.29254,12.573465",
		"Армения": "41.29254,12.573465",
		"Византия": "41.29254,12.573465",
		"Фергана": "41.29254,12.573465",
		"Иран": "41.29254,12.573465",
		"Коринф": "41.29254,12.573465",
		"Ирак, Антиохи, Магриб": "41.29254,12.573465",
		"Хузистан": "41.29254,12.573465",
		"Индия": "41.29254,12.573465",
		"Армения, Азербайджан": "41.29254,12.573465",
		"Ирак, Иран": "41.29254,12.573465",
		"Планета": "41.29254,12.573465",
		"Ирак": "41.29254,12.573465",
		"Египет, Сирия, Палестина": "41.29254,12.573465",
		"Армения, Турция": "41.29254,12.573465",
		"Ближний Восток": "41.29254,12.573465",
		"Палестина, Египет": "41.29254,12.573465",
		"Кавказ": "41.29254,12.573465",
		"Армения, Иран": "41.29254,12.573465",
		"Сицилия": "41.29254,12.573465",
		"Египет, Малая Азия": "41.29254,12.573465",
		"Малая Азия, Египет": "41.29254,12.573465",
		"Турция, Ирак": "41.29254,12.573465",
		"Малая Азия, Силиджия": "41.29254,12.573465",
		"Адриатическое \nморе": "41.29254,12.573465",
		"О. Крит": "41.29254,12.573465",
		"Рим": "41.29254,12.573465",
		"Египет, Сирия": "41.29254,12.573465",
		"Сирия, Ливия": "41.29254,12.573465",
		"Китай, Сирия, Ливия": "41.29254,12.573465",
		"Йемен": "41.29254,12.573465",
		"Афганистан": "41.29254,12.573465",
		"Германия": "41.29254,12.573465",
		"Словакия Югославия": "41.29254,12.573465",
		"Португалия": "41.29254,12.573465",
		"Венгрия": "41.29254,12.573465",
		"Венесуэла": "41.29254,12.573465",
		"Чили": "41.29254,12.573465",
		"Перу": "41.29254,12.573465",
		"Россия": "41.29254,12.573465",
		"Перу, Чили": "41.29254,12.573465",
		"Филиппины": "41.29254,12.573465",
		"Италия, Греция": "41.29254,12.573465",
		"Анатолия": "41.29254,12.573465",
		"Азербайджан, \nИран": "41.29254,12.573465",
		"Азербайджан": "41.29254,12.573465",
		"Китай или Корея": "41.29254,12.573465",
		"Индонезия": "41.29254,12.573465",
		"Эгейское море": "41.29254,12.573465",
		"Ямайка": "41.29254,12.573465",
		"Сицилия, Мальта": "41.29254,12.573465",
		"Алжир": "41.29254,12.573465",
		"Средний Восток": "41.29254,12.573465",
		"Антильские острова": "41.29254,12.573465",
		"Эквадор": "41.29254,12.573465",
		"Европа – \nСев. Африка": "41.29254,12.573465",
		"Ливан – Сирия": "41.29254,12.573465",
		"Острова Рюкю": "41.29254,12.573465",
		"Гватемала": "41.29254,12.573465",
		"Тайвань": "41.29254,12.573465",
		"Европа": "41.29254,12.573465",
		"Мексика": "41.29254,12.573465",
		"Крит": "41.29254,12.573465",
		"США": "41.29254,12.573465",
		"о. Бали": "41.29254,12.573465",
		"Индия, Пакистан": "41.29254,12.573465",
		"Испания": "41.29254,12.573465",
		"Молдавия": "41.29254,12.573465",
		"Грузия, Азербайджан": "41.29254,12.573465",
		"Пиренеи": "41.29254,12.573465",
		"Израиль, Сирия, Ливан": "41.29254,12.573465",
		"Турция, Кавказ": "41.29254,12.573465",
		"Перу, Чили, Аргентина": "41.29254,12.573465",
		"Гавайские острова": "41.29254,12.573465",
		"Перу, Чили, Боливия": "41.29254,12.573465",
		"Центральная Америка": "41.29254,12.573465",
		"Колумбия, Венесуэла": "41.29254,12.573465",
		"Средняя Азия": "41.29254,12.573465",
		"Гималаи": "41.29254,12.573465",
		"Иран, Туркмения": "41.29254,12.573465",
		"Туркмения": "41.29254,12.573465",
		"Гватемала, Сальвадор": "41.29254,12.573465",
		"Узбекистан": "41.29254,12.573465",
		"Турция, Армения": "41.29254,12.573465",
		"Эквадор, Колумбия": "41.29254,12.573465",
		"Калифорния": "41.29254,12.573465",
		"Таджикистан": "41.29254,12.573465",
		"Индонезия, Бали": "41.29254,12.573465",
		"Индия – Пакистан": "41.29254,12.573465",
		"Пакистан": "41.29254,12.573465",
		"Аргентина": "41.29254,12.573465",
		"Алеутские острова": "41.29254,12.573465",
		"Туркмения, Иран": "41.29254,12.573465",
		"Камчатка": "41.29254,12.573465",
		"Марокко": "41.29254,12.573465",
		"Аляска": "41.29254,12.573465",
		"Гватемала, Гондурас": "41.29254,12.573465",
		"Никарагуа": "41.29254,12.573465",
		"Италия, Австрия, Словения, Югославия": "41.29254,12.573465",
		"Бухарест": "41.29254,12.573465",
		"Иран, Восток": "41.29254,12.573465",
		"Австралия": "41.29254,12.573465",
		"Колумбия": "41.29254,12.573465",
		"Сальвадор": "41.29254,12.573465",
		"Новая Зеландия": "41.29254,12.573465",
		"о. Сахалин": "41.29254,12.573465",
		"Юго-Восточная \nАзия": "41.29254,12.573465",
		"Гаити": "41.29254,12.573465",
		"Месопотамия": "41.29254,12.573465",
		"Англия": "41.29254,12.573465",
		"Ирландия": "41.29254,12.573465",
		"Голландия": "41.29254,12.573465",
		"Черное море": "41.29254,12.573465",
		"Средиземное море (?)": "41.29254,12.573465",
		"Западная Европа": "41.29254,12.573465",
		"Русь": "41.29254,12.573465",
		"Атлантика": "41.29254,12.573465",
		"Междуречье": "41.29254,12.573465",
		"Голландия, Англия": "41.29254,12.573465",
		"Константинополь": "41.29254,12.573465",
		"Серное море": "41.29254,12.573465",
		"Северное море": "41.29254,12.573465",
		"Нидерланды": "41.29254,12.573465",
		"Польша, русские земли": "41.29254,12.573465",
		"Японское море": "41.29254,12.573465",
		"Финский залив": "41.29254,12.573465",
		"Австрия, Венгрия": "41.29254,12.573465",
		"Львов": "41.29254,12.573465",
		"О. Суматра": "41.29254,12.573465",
		"Фландрия, Зеландия, Голландия": "41.29254,12.573465",
		"Пуэрто-Рико": "41.29254,12.573465",
		"Франция": "41.29254,12.573465",
		"Мексиканский залив": "41.29254,12.573465",
		"Германия, Дания": "41.29254,12.573465",
		"Гваделупа": "41.29254,12.573465",
		"Остров Барбадос": "41.29254,12.573465",
		"Гибралтар": "41.29254,12.573465",
		"Балтийское море": "41.29254,12.573465",
		"Атлантика, \nСеверная Америка": "41.29254,12.573465",
		"Германия, Голландия": "41.29254,12.573465",
		"Амстердам": "41.29254,12.573465",
		"Великобритания": "41.29254,12.573465",
		"Уругвай": "41.29254,12.573465",
		"Флорида": "41.29254,12.573465",
		"Куба": "41.29254,12.573465",
		"Украина": "41.29254,12.573465",
		"Мартиника, Барбадос": "41.29254,12.573465",
		"Петербург": "41.29254,12.573465",
		"о. Ямайка": "41.29254,12.573465",
		"Карибское море": "41.29254,12.573465",
		"Польша": "41.29254,12.573465",
		"Гонконг": "41.29254,12.573465",
		"Мальта": "41.29254,12.573465",
		"Вьетнам": "41.29254,12.573465",
		"Каролинские острова": "41.29254,12.573465",
		"США, Пуэрто-Рико": "41.29254,12.573465",
		"о. Санто-Доминго": "41.29254,12.573465",
		"Бангладеш": "41.29254,12.573465",
		"Бенгалия": "41.29254,12.573465",
		"Канада": "41.29254,12.573465",
		"Мадагаскар": "41.29254,12.573465",
		"Гамбург": "41.29254,12.573465",
		"США, Канада": "41.29254,12.573465",
		"Гондурас": "41.29254,12.573465",
		"Тонга": "41.29254,12.573465",
		"Коморские острова": "41.29254,12.573465",
		"Перу, Эквадор": "41.29254,12.573465",
		"Таиланд": "41.29254,12.573465",
		"Грузия": "41.29254,12.573465",
		"Непал": "41.29254,12.573465",
		"Швейцария": "41.29254,12.573465",
		"ЮАР": "41.29254,12.573465",
		"Карибский бассейн": "41.29254,12.573465",
		"Австрия": "41.29254,12.573465",
		"Северная Корея": "41.29254,12.573465",
		"Чехия": "41.29254,12.573465",
		"Южная Америка": "41.29254,12.573465",
		"Гаити, Доминиканская республика": "41.29254,12.573465",
		"Доминиканская \nреспублика": "41.29254,12.573465",
		"США, \nКарибское море": "41.29254,12.573465",
		"Дальний Восток": "41.29254,12.573465",
		"Южная Европа": "41.29254,12.573465",
		"Аравия": "41.29254,12.573465",
		"Европа, Русь": "41.29254,12.573465",
		"Ирак, Малая Азия": "41.29254,12.573465",
		"Русь, Европа": "41.29254,12.573465",
		"Русь, Литва": "41.29254,12.573465",
		"Монголия": "41.29254,12.573465",
		"Швеция": "41.29254,12.573465",
		"Прибалтика": "41.29254,12.573465",
		"Европа, Китай": "41.29254,12.573465",
		"Восточная Европа": "41.29254,12.573465",
		"США, Европа": "41.29254,12.573465",
		"Бразилия": "41.29254,12.573465",
		"СССР": "41.29254,12.573465",
		"Индия, Эфиопия": "41.29254,12.573465",
		"Византия, \nЮг Руси": "41.29254,12.573465",
		"Аравийский полуостров": "41.29254,12.573465",
		"Адриатика": "41.29254,12.573465",
		"Европа, Россия, Китай": "41.29254,12.573465",
		"Московское государство": "41.29254,12.573465",
		"Скандинавия": "41.29254,12.573465",
		"Маньчжурия": "41.29254,12.573465",
		"Украина, Одесса, Николаев": "41.29254,12.573465",
		"Остров Мартиника": "41.29254,12.573465",
		"Исландия": "41.29254,12.573465",
		"Индонезия, Европа, Планета": "41.29254,12.573465",
		"Новая Гвинея": "41.29254,12.573465",
		"Оманский залив": "41.29254,12.573465",
		"Эгейское – Мраморное моря": "41.29254,12.573465",
		"Кипр": "41.29254,12.573465",
		"Персия": "41.29254,12.573465",
		"Пакистан, Индия": "41.29254,12.573465",
		"Македония": "41.29254,12.573465",
		"Сев. Восток Испании": "41.29254,12.573465",
		"Карфаген, Сев. Африка": "41.29254,12.573465",
		"Сибирь": "41.29254,12.573465",
		"Средиземноморье": "41.29254,12.573465",
		"Балканы": "41.29254,12.573465",
		"Западный Китай": "41.29254,12.573465",
		"Северо-Западный Китай": "41.29254,12.573465",
		"Галлия (Германия)": "41.29254,12.573465",
		"Юг Франции": "41.29254,12.573465",
		"Галлия": "41.29254,12.573465",
		"Северная Греция": "41.29254,12.573465",
		"Британия": "41.29254,12.573465",
		"Хорватия": "41.29254,12.573465",
		"Причерноморье": "41.29254,12.573465",
		"Фракия (Франция)": "41.29254,12.573465",
		"Восточный Рим": "41.29254,12.573465",
		"Франция, \nЗападный Рим": "41.29254,12.573465",
		"Северная Африка": "41.29254,12.573465",
		"Алтай": "41.29254,12.573465",
		"Сибирь, Средняя Азия": "41.29254,12.573465",
		"Пиренейский полуостров": "41.29254,12.573465",
		"Китай, Сибирь": "41.29254,12.573465",
		"Карпаты": "41.29254,12.573465",
		"Устье Волги": "41.29254,12.573465",
		"Западная Сибирь": "41.29254,12.573465",
		"Азия": "41.29254,12.573465",
		"Забайкалье": "41.29254,12.573465",
		"Северная Индия": "41.29254,12.573465",
		"Монголия, Китай": "41.29254,12.573465",
		"Средняя Азия, Корея": "41.29254,12.573465",
		"Север Кавказа": "41.29254,12.573465",
		"Украина, Крым": "41.29254,12.573465",
		"Средняя Азия, Монголия": "41.29254,12.573465",
		"Тибет": "41.29254,12.573465",
		"Малая Азия, Кавказ": "41.29254,12.573465",
		"Ирак, Грузия": "41.29254,12.573465",
		"Европа, Моравия": "41.29254,12.573465",
		"Индокитай": "41.29254,12.573465",
		"Юго-Восточная Азия": "41.29254,12.573465",
		"Шотландия": "41.29254,12.573465",
		"Франция, Нормандия": "41.29254,12.573465",
		"Золотая Орда": "41.29254,12.573465",
		"Балканский полуостров": "41.29254,12.573465",
		"Южная Русь": "41.29254,12.573465",
		"Болгария": "41.29254,12.573465",
		"Северная Америка": "41.29254,12.573465",
		"Венеция": "41.29254,12.573465",
		"Северная Италия": "41.29254,12.573465",
		"Северная Франция": "41.29254,12.573465",
		"Бельгия": "41.29254,12.573465",
		"Азорские острова": "41.29254,12.573465",
		"Корея": "41.29254,12.573465",
		"Ливония": "41.29254,12.573465",
		"Россия, Сибирь": "41.29254,12.573465",
		"Дания": "41.29254,12.573465",
		"Азов": "41.29254,12.573465",
		"Крым": "41.29254,12.573465",
		"Фландрия": "41.29254,12.573465",
		"Евразия": "41.29254,12.573465",
		"Маджент\n(Китай)": "41.29254,12.573465",
		"Южная Африка": "41.29254,12.573465",
		"Япония, Корея": "41.29254,12.573465",
		"Финляндия": "41.29254,12.573465",
		"Тихий океан": "41.29254,12.573465",
		"Европа, Индокитай": "41.29254,12.573465",
		"Филиппинское море": "41.29254,12.573465",
		"Тихий океан, Япония": "41.29254,12.573465",
		"Россия, Китай": "41.29254,12.573465",
		"Синайский полуостров": "41.29254,12.573465",
		"Конго": "41.29254,12.573465",
		"Западная Римская империя": "41.29254,12.573465",
		"Киев": "41.29254,12.573465",
		"Китай, Монголия": "41.29254,12.573465",
		"Камбоджа": "41.29254,12.573465",
		"Руанда": "41.29254,12.573465",
		"Древний мир": "41.29254,12.573465",
		"Древний Рим": "41.29254,12.573465",
		"Европа, Африка": "41.29254,12.573465",
		"Рим (империя)": "41.29254,12.573465",
		"Европа, Северная Африка, Азия, Аравия": "41.29254,12.573465",
		"Киевская Русь": "41.29254,12.573465",
		"Русь, Ливония": "41.29254,12.573465",
		"Кавказ, Азия, Орда": "41.29254,12.573465",
		"Европа, Ближний Восток, Индокитай": "41.29254,12.573465",
		"Карельская земля": "41.29254,12.573465",
		"Египет – Турция": "41.29254,12.573465",
		"Лондон": "41.29254,12.573465",
		"Европа, Малая Азия": "41.29254,12.573465",
		"Западнорусские земли": "41.29254,12.573465",
		"Москва, Вологда": "41.29254,12.573465",
		"Юг России": "41.29254,12.573465",
		"Юг Европы": "41.29254,12.573465",
		"Европа, Малая Азия, Египет": "41.29254,12.573465",
		"Европа, Малая Азия, Африка": "41.29254,12.573465",
		"Европа, Малая Азия, Африка, Индия": "41.29254,12.573465",
		"Бессарабия": "41.29254,12.573465",
		"Франция, Финляндия": "41.29254,12.573465",
		"Русская Америка": "41.29254,12.573465",
		"Африка, Китай": "41.29254,12.573465",
		"Месопотамия, Ливан": "41.29254,12.573465",
		"Панама": "41.29254,12.573465",
		"Китай, Индия": "41.29254,12.573465",
		"Эфиопия": "41.29254,12.573465",
		"Африка, Океания": "41.29254,12.573465",
		"Эстония": "41.29254,12.573465",
		"Латвия": "41.29254,12.573465",
		"Норвегия": "41.29254,12.573465",
		"Москва": "41.29254,12.573465",
		"Рига": "41.29254,12.573465",
		"Ливия": "41.29254,12.573465",
		"Альпы": "41.29254,12.573465",
		"Голгофа": "41.29254,12.573465",
		"Ближний и Средний Восток": "41.29254,12.573465",
		"Багдад": "41.29254,12.573465",
		"Литва, Польша": "41.29254,12.573465",
		"Орда": "41.29254,12.573465",
		"Мир (Европа)": "41.29254,12.573465",
		"Атлантика, Северная Америка": "41.29254,12.573465",
		"Южная Африка, Дикий берег": "41.29254,12.573465",
		"Бенгальский залив": "41.29254,12.573465",
		"Европа – США": "41.29254,12.573465",
		"Мир в целом": "41.29254,12.573465",
		"Арабский мир, возможно, \nмир в целом": "41.29254,12.573465",
		"Аттика": "41.29254,12.573465",
		"Карфаген": "41.29254,12.573465",
		"Римская империя": "41.29254,12.573465",
		"Европа, Ближний Восток": "41.29254,12.573465",
		"Румыния": "41.29254,12.573465",
		"Античный мир": "41.29254,12.573465",
		"Северная Русь": "41.29254,12.573465",
		"Индонезия\n(о. Ламбок)": "41.29254,12.573465",
		"Ливан": "41.29254,12.573465",
		"Великий Новгород": "41.29254,12.573465",
		"Англия, северная Франция": "41.29254,12.573465",
		"Боливия": "41.29254,12.573465",
		"Россия, Камчатка": "41.29254,12.573465",
		"Южный Урал": "41.29254,12.573465",
		"Восточная Африка": "41.29254,12.573465",
		"Западная Африка": "41.29254,12.573465",
		"Ява": "41.29254,12.573465",
		"Чехословакия": "41.29254,12.573465",
		"Восточное побережье \nТихого океана": "41.29254,12.573465",
		"Бангладеш, Индия, Непал": "41.29254,12.573465",
		"Фиджи": "41.29254,12.573465",
	}[reg]
}


func timeFormat(date disasterdb.Date) string {
	return fmt.Sprintf("%d/%d/%d", date.Day, date.Moth, date.Year)
}

type row struct {
	Titel string `json:"Titel"`
	Start string `json:"Start"`
	End string `json:"End"`
	Description string `json:"Description"`
	Image string `json:"Image"`
	Place string `json:"Place"`
	Location string `json:"Location"`
	Source string `json:"Source"`
	SourceURL string `json:"Source URL"`
}

func toRow(act disasterdb.Action) row {
	startDate := timeFormat(act.StartDate) 

	return row {
		Titel: "",
		Start: startDate,
		End: startDate, // for now, strange values in tables
		Description: "",
		Image: "",
		Place: "",
		Location: lookupLocation(act.Region),
		Source: "",
		SourceURL: "",
	}
}

func ToJson(_ context.Context, acts []disasterdb.Action) []byte {
	rsl := common.Fmap(func (act disasterdb.Action ) row { return toRow(act ) }, acts)

	res, err := json.Marshal(rsl)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
