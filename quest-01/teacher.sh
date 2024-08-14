#!/usr/bin/bash
export INTERVIEW_NUMBER=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2)
echo $INTERVIEW_NUMBER
cat "interviews/interview-$INTERVIEW_NUMBER"
echo $MAIN_SUSPECT
