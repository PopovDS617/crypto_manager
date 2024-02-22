MS_PRICE_RECEIVER=./app/ms_price_receiver
MS_PRICE_SAVER=./app/ms_price_saver


run-receiver:
	cd ${MS_PRICE_RECEIVER} && make run

run-saver:
	cd ${MS_PRICE_SAVER} && make run