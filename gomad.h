#include <stdio.h>
#include <stdint.h>
#include <limits.h>
#include <string.h>
#include <mad.h>

/* Size of the read buffer. */
#define BUFFER_SIZE (5 * 8192)

/*
 * MAD decoder wrapper structure.
 */
struct gomad_decoder {
	FILE *file;
	int current_sample;
	struct mad_stream stream;
	struct mad_frame frame;
	struct mad_header header;
	struct mad_synth synth;
	mad_timer_t timer;
	unsigned char buf[BUFFER_SIZE + MAD_BUFFER_GUARD];
};


/*
 * Initialize MAD decoder.
 * Returns -1 on failure.
 */
int gomad_open(struct gomad_decoder *decoder, const char *filename);



/*
 * Close MAD decoder and free all related resources.
 */
void gomad_close(struct gomad_decoder *decoder);
