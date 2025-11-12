package constants

const (
	QUERY_GET_TOTAL_ENVELOPES = "select * from public.get_total_envelopes_by_journal_promoter(i_journal_id => $1, i_promoter_id => $2);"
)
