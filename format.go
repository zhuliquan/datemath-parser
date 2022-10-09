package datemath_parser

const (
	EPOCH_MILLIS                            = "epoch_millis"
	EPOCH_SECOND                            = "epoch_second"
	DATE_OPTIONAL_TIME                      = "date_optional_time"
	STRICT_DATE_OPTIONAL_TIME               = "strict_date_optional_time"
	STRICT_DATE_OPTIONAL_TIME_NANOS         = "strict_date_optional_time_nanos"
	BASIC_DATE                              = "basic_date"
	BASIC_DATE_TIME                         = "basic_date_time"
	BASIC_DATE_TIME_NO_MILLIS               = "basic_date_time_no_millis"
	BASIC_ORDINAL_DATE                      = "basic_ordinal_date"
	BASIC_ORDINAL_DATE_TIME                 = "basic_ordinal_date_time"
	BASIC_ORDINAL_DATE_TIME_NO_MILLIS       = "basic_ordinal_date_time_no_millis"
	BASIC_TIME                              = "basic_time"
	BASIC_TIME_NO_MILLIS                    = "basic_time_no_millis"
	BASIC_T_TIME                            = "basic_t_time"
	BASIC_T_TIME_NO_MILLIS                  = "basic_t_time_no_millis"
	BASIC_WEEK_DATE                         = "basic_week_date"
	STRICT_BASIC_WEEK_DATE                  = "strict_basic_week_date"
	BASIC_WEEK_DATE_TIME                    = "basic_week_date_time"
	STRICT_BASIC_WEEK_DATE_TIME             = "strict_basic_week_date_time"
	BASIC_WEEK_DATE_TIME_NO_MILLIS          = "basic_week_date_time_no_millis"
	STRICT_BASIC_WEEK_DATE_TIME_NO_MILLIS   = "strict_basic_week_date_time_no_millis"
	DATE                                    = "date"
	STRICT_DATE                             = "strict_date"
	DATE_HOUR                               = "date_hour"
	STRICT_DATE_HOUR                        = "strict_date_hour"
	DATE_HOUR_MINUTE                        = "date_hour_minute"
	STRICT_DATE_HOUR_MINUTE                 = "strict_date_hour_minute"
	DATE_HOUR_MINUTE_SECOND                 = "date_hour_minute_second"
	STRICT_DATE_HOUR_MINUTE_SECOND          = "strict_date_hour_minute_second"
	DATE_HOUR_MINUTE_SECOND_FRACTION        = "date_hour_minute_second_fraction"
	STRICT_DATE_HOUR_MINUTE_SECOND_FRACTION = "strict_date_hour_minute_second_fraction"
	DATE_HOUR_MINUTE_SECOND_MILLIS          = "date_hour_minute_second_millis"
	STRICT_DATE_HOUR_MINUTE_SECOND_MILLIS   = "strict_date_hour_minute_second_millis"
	DATE_TIME                               = "date_time"
	STRICT_DATE_TIME                        = "strict_date_time"
	DATE_TIME_NO_MILLIS                     = "date_time_no_millis"
	HOUR                                    = "hour"
	TIME_NO_MILLIS                          = "time_no_millis"
	HOUR_MINUTE                             = "hour_minute"
	HOUR_MINUTE_SECOND                      = "hour_minute_second"
	HOUR_MINUTE_SECOND_FRACTION             = "hour_minute_second_fraction"
	HOUR_MINUTE_SECOND_MILLIS               = "hour_minute_second_millis"
	ORDINAL_DATE                            = "ordinal_date"
	ORDINAL_DATE_TIME                       = "ordinal_date_time"
	ORDINAL_DATE_TIME_NO_MILLIS             = "ordinal_date_time_no_millis"
	TIME                                    = "time"
	T_TIME                                  = "t_time"
	T_TIME_NO_MILLIS                        = "t_time_no_millis"
	WEEK_DATE                               = "week_date"
	WEEK_DATE_TIME                          = "week_date_time"
	WEEK_DATE_TIME_NO_MILLIS                = "week_date_time_no_millis"
	WEEKYEAR                                = "weekyear"
	WEEKYEAR_WEEK                           = "weekyear_week"
	WEEKYEAR_WEEK_DAY                       = "weekyear_week_day"
	YEAR_MONTH                              = "year_month"
	YEAR                                    = "year"
	YEAR_MONTH_DAY                          = "year_month_day"
	STRICT_TIME_NO_MILLIS                   = "strict_time_no_millis"
	STRICT_DATE_TIME_NO_MILLIS              = "strict_date_time_no_millis"
	STRICT_HOUR                             = "strict_hour"
	STRICT_HOUR_MINUTE                      = "strict_hour_minute"
	STRICT_HOUR_MINUTE_SECOND               = "strict_hour_minute_second"
	STRICT_HOUR_MINUTE_SECOND_FRACTION      = "strict_hour_minute_second_fraction"
	STRICT_HOUR_MINUTE_SECOND_MILLIS        = "strict_hour_minute_second_millis"
	STRICT_ORDINAL_DATE                     = "strict_ordinal_date"
	STRICT_ORDINAL_DATE_TIME                = "strict_ordinal_date_time"
	STRICT_ORDINAL_DATE_TIME_NO_MILLIS      = "strict_ordinal_date_time_no_millis"
	STRICT_TIME                             = "strict_time"
	STRICT_T_TIME                           = "strict_t_time"
	STRICT_T_TIME_NO_MILLIS                 = "strict_t_time_no_millis"
	STRICT_WEEK_DATE                        = "strict_week_date"
	STRICT_WEEK_DATE_TIME                   = "strict_week_date_time"
	STRICT_WEEK_DATE_TIME_NO_MILLIS         = "strict_week_date_time_no_millis"
	STRICT_WEEKYEAR                         = "strict_weekyear"
	STRICT_WEEKYEAR_WEEK                    = "strict_weekyear_week"
	STRICT_WEEKYEAR_WEEK_DAY                = "strict_weekyear_week_day"
	STRICT_YEAR_MONTH                       = "strict_year_month"
	STRICT_YEAR                             = "strict_year"
	STRICT_YEAR_MONTH_DAY                   = "strict_year_month_day"
)

// built-in format
var BuiltInFormat = map[string][]string{
	// A formatter for the number of milliseconds since the epoch.
	// Note, that this timestamp is subject to the limits of a Java Long.MIN_VALUE and Long.MAX_VALUE.
	EPOCH_MILLIS: {EPOCH_MILLIS},
	// A formatter for the number of seconds since the epoch.
	// Note, that this timestamp is subject to the limits of a Java Long.MIN_VALUE and Long.
	// MAX_VALUE divided by 1000 (the number of milliseconds in a second).
	EPOCH_SECOND: {EPOCH_SECOND},

	// A generic ISO datetime parser, where the date must include the year at a minimum, and the time (separated by T), is optional. Examples: yyyy-MM-dd'T'HH:mm:ss.SSSZ or yyyy-MM-dd.
	DATE_OPTIONAL_TIME:        {"yyyy-MM-ddTHH:mm:ss.SSSZ", "yyyy-MM-dd"},
	STRICT_DATE_OPTIONAL_TIME: {"yyyy-MM-ddTHH:mm:ss.SSSZ", "yyyy-MM-dd"},
	// A generic ISO datetime parser, where the date must include the year at a minimum, and the time (separated by T), is optional. The fraction of a second part has a nanosecond resolution. Examples: yyyy-MM-ddTHH:mm:ss.SSSSSSZ or yyyy-MM-dd.
	STRICT_DATE_OPTIONAL_TIME_NANOS: {"yyyy-MM-ddTHH:mm:ss.SSSSSSZ", "yyyy-MM-dd"},

	// A basic formatter for a full date as four digit year, two digit month of year, and two digit day of month: yyyyMMdd.
	BASIC_DATE: {"yyyyMMdd"},

	// A basic formatter that combines a basic date and time, separated by a T: .
	BASIC_DATE_TIME: {"yyyyMMddTHHmmss.SSSZ"},

	// A basic formatter that combines a basic date and time without millis, separated by a T: yyyyMMddTHHmmssZ.
	BASIC_DATE_TIME_NO_MILLIS: {"yyyyMMddTHHmmssZ"},

	// A formatter for a full ordinal date, using a four digit year and three digit dayOfYear: yyyyDDD.
	BASIC_ORDINAL_DATE: {"yyyyDDD"},

	// A formatter for a full ordinal date and time, using a four digit year and three digit dayOfYear: yyyyDDDTHHmmss.SSSZ.
	BASIC_ORDINAL_DATE_TIME: {"yyyyDDDTHHmmss.SSSZ"},

	// A formatter for a full ordinal date and time without millis, using a four digit year and three digit dayOfYear: yyyyDDDTHHmmssZ.
	BASIC_ORDINAL_DATE_TIME_NO_MILLIS: {"yyyyDDDTHHmmssZ"},

	// A basic formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, three digit millis, and time zone offset: HHmmss.SSSZ.
	BASIC_TIME: {"HHmmss.SSSZ"},

	// A basic formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, and time zone offset: HHmmssZ.
	BASIC_TIME_NO_MILLIS: {"HHmmssZ"},

	// A basic formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, three digit millis, and time zone off set prefixed by T: THHmmss.SSSZ.
	BASIC_T_TIME: {"THHmmss.SSSZ"},

	// A basic formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, and time zone offset prefixed by T: THHmmssZ.
	BASIC_T_TIME_NO_MILLIS: {"THHmmssZ"},

	// A basic formatter for a full date as four digit weekyear, two digit week of weekyear, and one digit day of week: xxxxWwwe.
	BASIC_WEEK_DATE:        {"xxxxWwwe"},
	STRICT_BASIC_WEEK_DATE: {"xxxxWwwe"},

	// A basic formatter that combines a basic weekyear date and time, separated by a T: xxxxWwweTHHmmss.SSSZ.
	BASIC_WEEK_DATE_TIME:        {"xxxxWwweTHHmmss.SSSZ"},
	STRICT_BASIC_WEEK_DATE_TIME: {"xxxxWwweTHHmmss.SSSZ"},

	// A basic formatter that combines a basic weekyear date and time without millis, separated by a T: xxxxWwweTHHmmssZ.
	BASIC_WEEK_DATE_TIME_NO_MILLIS:        {"xxxxWwweTHHmmssZ"},
	STRICT_BASIC_WEEK_DATE_TIME_NO_MILLIS: {"xxxxWwweTHHmmssZ"},

	// A formatter for a full date as four digit year, two digit month of year, and two digit day of month: yyyy-MM-dd.
	DATE:        {"yyyy-MM-dd"},
	STRICT_DATE: {"yyyy-MM-dd"},

	// A formatter that combines a full date and two digit hour of day: yyyy-MM-ddTHH.
	DATE_HOUR:        {"yyyy-MM-ddTHH"},
	STRICT_DATE_HOUR: {"yyyy-MM-ddTHH"},

	// A formatter that combines a full date, two digit hour of day, and two digit minute of hour: yyyy-MM-ddTHH:mm.
	DATE_HOUR_MINUTE:        {"yyyy-MM-ddTHH:mm"},
	STRICT_DATE_HOUR_MINUTE: {"yyyy-MM-ddTHH:mm"},

	// A formatter that combines a full date, two digit hour of day, two digit minute of hour, and two digit second of minute: yyyy-MM-ddTHH:mm:ss.
	DATE_HOUR_MINUTE_SECOND:        {"yyyy-MM-ddTHH:mm:ss"},
	STRICT_DATE_HOUR_MINUTE_SECOND: {"yyyy-MM-ddTHH:mm:ss"},

	// A formatter that combines a full date, two digit hour of day, two digit minute of hour, two digit second of minute, and three digit fraction of second: yyyy-MM-ddTHH:mm:ss.SSS.
	DATE_HOUR_MINUTE_SECOND_FRACTION:        {"yyyy-MM-ddTHH:mm:ss.SSS"},
	STRICT_DATE_HOUR_MINUTE_SECOND_FRACTION: {"yyyy-MM-ddTHH:mm:ss.SSS"},

	// A formatter that combines a full date, two digit hour of day, two digit minute of hour, two digit second of minute, and three digit fraction of second: yyyy-MM-ddTHH:mm:ss.SSS.
	DATE_HOUR_MINUTE_SECOND_MILLIS:        {"yyyy-MM-ddTHH:mm:ss.SSS"},
	STRICT_DATE_HOUR_MINUTE_SECOND_MILLIS: {"yyyy-MM-ddTHH:mm:ss.SSS"},

	// A formatter that combines a full date and time, separated by a T: yyyy-MM-ddTHH:mm:ss.SSSZ.
	DATE_TIME:        {"yyyy-MM-ddTHH:mm:ss.SSSZ"},
	STRICT_DATE_TIME: {"yyyy-MM-ddTHH:mm:ss.SSSZ"},

	// A formatter that combines a full date and time without millis, separated by a T: yyyy-MM-ddTHH:mm:ssZ.
	DATE_TIME_NO_MILLIS:        {"yyyy-MM-ddTHH:mm:ssZ"},
	STRICT_DATE_TIME_NO_MILLIS: {"yyyy-MM-ddTHH:mm:ssZ"},

	// A formatter for a two digit hour of day: HH
	HOUR:        {"HH"},
	STRICT_HOUR: {"HH"},

	// A formatter for a two digit hour of day and two digit minute of hour: HH:mm.
	HOUR_MINUTE:        {"HH:mm"},
	STRICT_HOUR_MINUTE: {"HH:mm"},

	// A formatter for a two digit hour of day, two digit minute of hour, and two digit second of minute: HH:mm:ss.
	HOUR_MINUTE_SECOND:        {"HH:mm:ss"},
	STRICT_HOUR_MINUTE_SECOND: {"HH:mm:ss"},

	// A formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, and three digit fraction of second: HH:mm:ss.SSS.
	HOUR_MINUTE_SECOND_FRACTION:        {"HH:mm:ss.SSS"},
	STRICT_HOUR_MINUTE_SECOND_FRACTION: {"HH:mm:ss.SSS"},

	// A formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, and three digit fraction of second: HH:mm:ss.SSS.
	HOUR_MINUTE_SECOND_MILLIS:        {"HH:mm:ss.SSS"},
	STRICT_HOUR_MINUTE_SECOND_MILLIS: {"HH:mm:ss.SSS"},

	// A formatter for a full ordinal date, using a four digit year and three digit dayOfYear: yyyy-DDD.
	ORDINAL_DATE:        {"yyyy-DDD"},
	STRICT_ORDINAL_DATE: {"yyyy-DDD"},

	// A formatter for a full ordinal date and time, using a four digit year and three digit dayOfYear: yyyy-DDDTHH:mm:ss.SSSZ.
	ORDINAL_DATE_TIME:        {"yyyy-DDDTHH:mm:ss.SSSZ"},
	STRICT_ORDINAL_DATE_TIME: {"yyyy-DDDTHH:mm:ss.SSSZ"},

	// A formatter for a full ordinal date and time without millis, using a four digit year and three digit dayOfYear: yyyy-DDDTHH:mm:ssZ.
	ORDINAL_DATE_TIME_NO_MILLIS:        {"yyyy-DDDTHH:mm:ssZ"},
	STRICT_ORDINAL_DATE_TIME_NO_MILLIS: {"yyyy-DDDTHH:mm:ssZ"},

	// A formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, three digit fraction of second, and time zone offset: HH:mm:ss.SSSZ.
	TIME:        {"HH:mm:ss.SSSZ"},
	STRICT_TIME: {"HH:mm:ss.SSSZ"},

	// A formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, and time zone offset: HH:mm:ssZ.
	TIME_NO_MILLIS:        {"HH:mm:ssZ"},
	STRICT_TIME_NO_MILLIS: {"HH:mm:ssZ"},

	// A formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, three digit fraction of second, and time zone offset prefixed by T: THH:mm:ss.SSSZ.
	T_TIME:        {"THH:mm:ss.SSSZ"},
	STRICT_T_TIME: {"THH:mm:ss.SSSZ"},

	// A formatter for a two digit hour of day, two digit minute of hour, two digit second of minute, and time zone offset prefixed by T: THH:mm:ssZ.
	T_TIME_NO_MILLIS:        {"THH:mm:ssZ"},
	STRICT_T_TIME_NO_MILLIS: {"THH:mm:ssZ"},

	// A formatter for a full date as four digit weekyear, two digit week of weekyear, and one digit day of week: xxxx-Www-e.
	WEEK_DATE:        {"xxxx-Www-e"},
	STRICT_WEEK_DATE: {"xxxx-Www-e"},

	// A formatter that combines a full weekyear date and time, separated by a T: xxxx-Www-eTHH:mm:ss.SSSZ.
	WEEK_DATE_TIME:        {"xxxx-Www-eTHH:mm:ss.SSSZ"},
	STRICT_WEEK_DATE_TIME: {"xxxx-Www-eTHH:mm:ss.SSSZ"},

	// A formatter that combines a full weekyear date and time without millis, separated by a T: xxxx-Www-eTHH:mm:ssZ.
	WEEK_DATE_TIME_NO_MILLIS:        {"xxxx-Www-eTHH:mm:ssZ"},
	STRICT_WEEK_DATE_TIME_NO_MILLIS: {"xxxx-Www-eTHH:mm:ssZ"},

	// A formatter for a four digit weekyear: xxxx.
	WEEKYEAR:        {"xxxx"},
	STRICT_WEEKYEAR: {"xxxx"},

	// A formatter for a four digit weekyear and two digit week of weekyear: xxxx-Www.
	WEEKYEAR_WEEK:        {"xxxx-Www"},
	STRICT_WEEKYEAR_WEEK: {"xxxx-Www"},

	// A formatter for a four digit weekyear, two digit week of weekyear, and one digit day of week: xxxx-Www-e.
	WEEKYEAR_WEEK_DAY:        {"xxxx-Www-e"},
	STRICT_WEEKYEAR_WEEK_DAY: {"xxxx-Www-e"},

	// A formatter for a four digit year and two digit month of year: yyyy-MM.
	YEAR_MONTH:        {"yyyy-MM"},
	STRICT_YEAR_MONTH: {"yyyy-MM"},

	// A formatter for a four digit year: yyyy.
	YEAR:        {"yyyy"},
	STRICT_YEAR: {"yyyy"},

	// A formatter for a four digit year, two digit month of year, and two digit day of month: yyyy-MM-dd.
	YEAR_MONTH_DAY:        {"yyyy-MM-dd"},
	STRICT_YEAR_MONTH_DAY: {"yyyy-MM-dd"},
}
