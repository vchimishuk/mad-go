#include "gomad.h"


int gomad_open(struct gomad_decoder *decoder, const char *filename)
{
	decoder->current_sample = 0;

	decoder->file = fopen(filename, "r");
	if (decoder->file == NULL) {
		return -1;
	}
	
	mad_stream_init(&decoder->stream);
	mad_frame_init(&decoder->frame);
	mad_header_init(&decoder->header);
	mad_synth_init(&decoder->synth);
	mad_timer_reset(&decoder->timer);
	
	return 0;
}


void gomad_close(struct gomad_decoder *decoder)
{
	/*
	mad_frame_finish(&decoder->frame);
	mad_stream_finish(&decoder->stream);
	mad_synth_finish(&decoder->synth);

	fclose(decoder->file);
	*/
}
